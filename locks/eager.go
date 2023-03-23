package locks

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrEagerLockFailed = errors.New("eager lock: failed to acquire lock")
)

type EagerLock struct {
	retries  int
	interval time.Duration
	register struct {
		sync.RWMutex
		m map[string]int64
	}
}

func NewEager() *EagerLock {
	return &EagerLock{
		retries:  3,
		interval: 5 * time.Second,
		register: struct {
			sync.RWMutex
			m map[string]int64
		}{m: make(map[string]int64)},
	}
}

func (e *EagerLock) LockWithRetries(key string, value int64) error {
	for i := 0; i <= e.retries; i++ {
		err := e.Lock(key, value)
		if err == nil {
			return nil
		}

		time.Sleep(e.interval)
	}
	return ErrEagerLockFailed
}

func (e *EagerLock) Lock(key string, value int64) error {
	timeout, exist := e.readLock(key)
	if !exist || time.Now().UnixNano() > timeout {
		e.register.Lock()
		defer e.register.Unlock()
		e.register.m[key] = value
		return nil
	}
	return ErrEagerLockFailed
}

func (e *EagerLock) readLock(key string) (int64, bool) {
	e.register.RLock()
	defer e.register.RUnlock()
	timeout, exist := e.register.m[key]
	return timeout, exist
}
