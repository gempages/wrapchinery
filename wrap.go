package wrapchinery

import (
	"context"
	"reflect"
	"runtime"
	"time"

	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	"github.com/RichardKnop/machinery/v2/backends/result"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	"github.com/RichardKnop/machinery/v2/config"
	lockiface "github.com/RichardKnop/machinery/v2/locks/iface"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gempages/wrapchinery/logger"
	"github.com/google/uuid"
)

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

// WrapNewWorker creates a new machinery worker with a random UUID as tag and concurrency = number of CPU x2
func (m *Server) WrapNewWorker() *machinery.Worker {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return m.NewWorker(uid.String(), runtime.NumCPU()*2)
}

// WrapSendTask calls machinery's SendTask function with task signature created using GetTaskSignature function
func (m *Server) WrapSendTask(taskName string, delay time.Duration, retry int, args ...interface{}) (*result.AsyncResult, error) {
	task := GetTaskSignature(taskName, delay, retry, args...)
	return m.SendTask(task)
}

func (m *Server) WrapSendTaskWithContext(
	ctx context.Context, taskName string, delay time.Duration, retry int, args ...interface{},
) (*result.AsyncResult, error) {
	task := GetTaskSignature(taskName, delay, retry, args...)
	return m.SendTaskWithContext(ctx, task)
}

// GetTaskSignature returns machinery's task signature object to use with SendTask and SendTaskWithContext functions
func GetTaskSignature(taskName string, delay time.Duration, retry int, args ...interface{}) *tasks.Signature {
	task := tasks.Signature{
		Name: taskName,
		Args: []tasks.Arg{},
	}
	if delay > 0 {
		timeETA := time.Now().UTC().Add(delay)
		task.ETA = &timeETA
	}
	task.RetryCount = retry
	task.IgnoreWhenTaskNotRegistered = true
	if len(args) > 0 {
		task.Args = parseArgs(args...)
	}
	return &task
}

func parseArgs(args ...interface{}) []tasks.Arg {
	taskArgs := []tasks.Arg{}
	for k := range args {
		taskArgs = append(taskArgs, tasks.Arg{
			Type:  reflect.TypeOf(args[k]).String(),
			Value: args[k],
		})
	}
	return taskArgs
}
