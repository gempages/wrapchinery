package wrapchinery

import (
	"context"
	"fmt"

	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	"github.com/RichardKnop/machinery/v2/backends/result"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	taskconfig "github.com/RichardKnop/machinery/v2/config"
	lockiface "github.com/RichardKnop/machinery/v2/locks/iface"
	redislock "github.com/RichardKnop/machinery/v2/locks/redis"
	"github.com/alicebob/miniredis"
	"github.com/gempages/go-helper/cache"
	"github.com/google/uuid"
)

var server *TaskServer

type TaskServer struct {
	machinery.Server
}

// WrapNewWorker creates a new machinery worker with a random UUID as tag
func (m *TaskServer) WrapNewWorker(concurrency int) *machinery.Worker {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return m.NewWorker(uid.String(), concurrency)
}

// WrapSendTask calls machinery's SendTask function with task signature created using GetTaskSignature function
func (m *TaskServer) WrapSendTask(cfg *TaskConfig, args ...interface{}) (*result.AsyncResult, error) {
	task := GetTaskSignature(cfg, args...)
	return m.SendTask(task)
}

func (m *TaskServer) WrapSendTaskWithContext(ctx context.Context, cfg *TaskConfig, args ...interface{}) (*result.AsyncResult, error) {
	task := GetTaskSignature(cfg, args...)
	return m.SendTaskWithContext(ctx, task)
}

// NewServer creates Server instance
func newServer(
	cnf *config.Config, brokerServer brokersiface.Broker,
	backendServer backendsiface.Backend, lock lockiface.Lock,
) *TaskServer {

	return &TaskServer{
		*machinery.NewServer(cnf, brokerServer, backendServer, lock),
	}
}

// InitTaskServer initializes task server. Call this before registering tasks for workers.
func InitTaskServer() error {
	cnf, err := taskconfig.NewFromEnvironment()
	if err != nil {
		return err
	}

	broker := redisbroker.NewGR(cnf, []string{cnf.Broker}, 0)
	backend := redisbackend.NewGR(cnf, []string{cnf.ResultBackend}, 0)
	lock := redislock.New(cnf, []string{cnf.Broker}, 0, 5)
	server = newServer(cnf, broker, backend, lock)
	return nil
}

// InitTestTaskServer initializes task server for automation tests. Call this before registering tasks for workers.
func InitTestTaskServer() error {
	cnf, err := taskconfig.NewFromEnvironment()
	if err != nil {
		return err
	}

	// Open mock redis
	mr, err := miniredis.Run()
	if err != nil {
		return err
	}

	err = cache.InitRedisClient(mr.Host(), mr.Port())
	if err != nil {
		return err
	}

	mockRedisAddress := fmt.Sprintf("%s:%s", mr.Host(), mr.Port())
	cnf.Broker = mockRedisAddress
	cnf.ResultBackend = mockRedisAddress
	broker := redisbroker.NewGR(cnf, []string{cnf.Broker}, 0)
	backend := redisbackend.NewGR(cnf, []string{cnf.ResultBackend}, 0)
	lock := redislock.New(cnf, []string{cnf.Broker}, 0, 5)
	server = newServer(cnf, broker, backend, lock)
	return nil
}

func Server() *TaskServer {
	return server
}
