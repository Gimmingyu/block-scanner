package common

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Mapper interface {
	TryGet(key string) (value interface{}, ok bool)
	Set(key string, value interface{})
	TryDel(key string) (ok bool)
	Del(key string)
}

type LifeCycle interface {
	BeforeStart(...func(c context.Context) error) error
	AfterStart()
}

type ctx struct {
	keys map[any]any

	cause error
	err   error
	mu    sync.Mutex
	done  atomic.Value
}

func (c *ctx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *ctx) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}

func (c *ctx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *ctx) Value(key any) any {
	return c.keys[key]
}

func New() context.Context {
	return &ctx{
		keys: make(map[any]any),
	}
}
