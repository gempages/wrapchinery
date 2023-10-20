package wrapchinery

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	"github.com/RichardKnop/machinery/v2/backends/result"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	"github.com/RichardKnop/machinery/v2/config"
	lockiface "github.com/RichardKnop/machinery/v2/locks/iface"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	gplog "github.com/gempages/go-helper/log"
	"github.com/gempages/go-helper/tracing"
	"github.com/gempages/wrapchinery/logger"
	"github.com/getsentry/sentry-go"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
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

type Server struct {
	machinery.Server
}

// NewServer creates Server instance
func NewServer(
	cnf *config.Config, brokerServer brokersiface.Broker,
	backendServer backendsiface.Backend, lock lockiface.Lock,
) *Server {

	server := &Server{
		*machinery.NewServer(cnf, brokerServer, backendServer, lock),
	}

	return server
}

func SetupLoggers() {
	log.SetDebug(logger.NewDebugLogger())
	log.SetError(logger.NewErrorLogger())
	log.SetInfo(logger.NewInfoLogger())
	log.SetFatal(logger.NewFatalLogger())
	log.SetWarning(logger.NewWarningLogger())
}

// WrapNewWorker creates a new worker worker with a random UUID as tag
func (m *Server) WrapNewWorker(concurrency int) *machinery.Worker {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	w := m.NewWorker(uid.String(), concurrency)
	// Override default error handler function in order to remove task ID in error message,
	// which makes it easier to manage issues in Sentry
	w.SetErrorHandler(func(err error) {
		gplog.Error(context.Background(), fmt.Errorf("worker task failed: %w", err))
	})
	return w
}

// WrapSendTask calls worker's SendTask function with task signature created using GetTaskSignature function
func (m *Server) WrapSendTask(cfg *TaskConfig, args ...interface{}) (*result.AsyncResult, error) {
	task := GetTaskSignature(cfg, args...)
	return m.SendTask(task)
}

func (m *Server) WrapSendTaskWithContext(ctx context.Context, cfg *TaskConfig, args ...interface{}) (*result.AsyncResult, error) {
	var err error

	span := sentry.StartSpan(ctx, "send_task")
	span.Description = cfg.Name
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

	task := GetTaskSignature(cfg, args...)
	return m.SendTaskWithContext(ctx, task)
}

// GetTaskSignature returns worker's task signature object to use with SendTask and SendTaskWithContext functions
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
