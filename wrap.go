package wrapchinery

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	gplog "github.com/gempages/go-helper/log"
	"github.com/gempages/go-helper/tracing"
	"github.com/gempages/wrapchinery/logger"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cast"
)

type TaskConfig struct {
	Name       string
	ShopID     uint64
	Delay      time.Duration
	RetryCount int
	OnSuccess  *TaskConfig
	OnError    *TaskConfig
}

const ShopIDHeader = "shopID"

func SetupLoggers() {
	log.SetDebug(logger.NewDebugLogger())
	log.SetError(logger.NewErrorLogger())
	log.SetInfo(logger.NewInfoLogger())
	log.SetFatal(logger.NewFatalLogger())
	log.SetWarning(logger.NewWarningLogger())
}

// GetTaskSignature returns machinery's task signature object to use with SendTask and SendTaskWithContext functions
func GetTaskSignature(cfg *TaskConfig, args ...interface{}) *tasks.Signature {
	task, _ := tasks.NewSignature(cfg.Name, parseArgs(args...))
	if cfg.Delay > 0 {
		timeETA := time.Now().UTC().Add(cfg.Delay)
		task.ETA = &timeETA
	}
	// manual add shopID into header
	if cfg.ShopID > 0 {
		if task.Headers == nil {
			task.Headers = tasks.Headers{}
		}
		task.Headers.Set(ShopIDHeader, cast.ToString(cfg.ShopID))
	}

	task.RetryCount = cfg.RetryCount
	task.IgnoreWhenTaskNotRegistered = true
	if cfg.OnSuccess != nil {
		task.OnSuccess = []*tasks.Signature{GetTaskSignature(cfg.OnSuccess, args...)}
	}
	if cfg.OnError != nil {
		task.OnError = []*tasks.Signature{GetTaskSignature(cfg.OnError, args...)}
	}
	return task
}

func RegisterTasks(taskList map[string]interface{}) error {
	return server.RegisterTasks(taskList)
}

func SendTask(ctx context.Context, task *TaskConfig, args ...interface{}) (*result.AsyncResult, error) {
	var err error

	span := sentry.StartSpan(ctx, "send_task")
	span.Description = task.Name
	defer func() {
		tracing.FinishSpan(span, err)
	}()
	ctx = span.Context()
	traceEncoded, err := json.Marshal(tracing.ToSentryTrace(span))
	if err != nil {
		gplog.Warning(ctx, fmt.Errorf("encode trace header: %w", err))
	} else {
		args = append([]interface{}{traceEncoded}, args...)
	}

	asyncResult, err := server.WrapSendTaskWithContext(ctx, task, args...)
	if err != nil {
		return asyncResult, fmt.Errorf("send task: %w", err)
	}
	return asyncResult, err
}

func SendTaskWaitResult(ctx context.Context, task *TaskConfig, args ...interface{}) error {
	var err error

	span := sentry.StartSpan(ctx, "send_task_wait")
	span.Description = task.Name
	defer func() {
		tracing.FinishSpan(span, err)
	}()
	ctx = span.Context()
	traceEncoded, err := json.Marshal(tracing.ToSentryTrace(span))
	if err != nil {
		gplog.Warning(ctx, fmt.Errorf("encode trace header: %w", err))
	} else {
		args = append([]interface{}{traceEncoded}, args...)
	}

	asyncResult, err := server.WrapSendTaskWithContext(ctx, task, args...)
	if err != nil {
		return fmt.Errorf("send task: %w", err)
	}

	_, err = asyncResult.Get(time.Millisecond * 5)

	if err != nil {
		return fmt.Errorf("get task result: %w", err)
	}
	return nil
}

func parseArgs(args ...interface{}) []tasks.Arg {
	var taskArgs []tasks.Arg
	for k := range args {
		taskArgs = append(taskArgs, tasks.Arg{
			Type:  reflect.TypeOf(args[k]).String(),
			Value: args[k],
		})
	}
	return taskArgs
}
