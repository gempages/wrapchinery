# wrapchinery
Wrapped machinery package (v2) to send tasks more easily

## Usage
Call `wrapchinery.NewServer()` instead of `machinery.NewServer()`
```go
import (
    github.com/es-hs/wrapchinery
)

server := wrapchinery.NewServer(...)
```

Helper function to ease the pain of creating Signature:
```go
GetTaskSignature(taskName string, delay time.Duration, retry int, args ...interface{}) *tasks.Signature
```