# wrapchinery

Wrapped machinery package (v2) to send tasks more easily

## Usage

Call `wrapchinery.NewServer()` instead of `machinery.NewServer()`

```go
import (
github.com/gempages/wrapchinery
)

server := wrapchinery.NewServer(...)
```

Use wrapped functions:

```go
// Create new worker
server.WrapNewWorker(concurrency int) *machinery.Worker
// Send a task
server.WrapSendTask(taskName string, delay time.Duration, retry int, args ...interface{}) (*result.AsyncResult, error)
// Send a task with context
server.WrapSendTaskWithContext(taskName string, ctx context.Context, delay time.Duration, retry int, args ...interface{}) (*result.AsyncResult, error)
```

Helper function to ease the pain of creating Signature:

```go
wrapchinery.GetTaskSignature(taskName string, delay time.Duration, retry int, args ...interface{}) *tasks.Signature
```
