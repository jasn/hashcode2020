package main

import (
	"context"
	"sync"

	"go.uber.org/multierr"
)

// ThrottledExecutor is a pooled executor which ensures that only a specified number of goroutines run concurrently.
// Errors do not block execution, but they are collected into a single multierr.
// Execution stops if the context times out.
type ThrottledExecutor struct {
	ctx           context.Context
	wg            sync.WaitGroup
	lock          sync.Mutex
	cancelledOnce sync.Once
	concurrency   chan struct{}

	err error
}

// Option is an option for the throttled executor.
type Option func(executor *ThrottledExecutor)

// NewThrottledExecutor creates a new executor.
func NewThrottledExecutor(ctx context.Context, concurrency int, options ...Option) *ThrottledExecutor {
	e := &ThrottledExecutor{
		ctx:         ctx,
		concurrency: make(chan struct{}, concurrency),
	}

	for _, opt := range options {
		opt(e)
	}

	return e
}

// Go adds a function to be executed.
func (t *ThrottledExecutor) Go(f func() error) {
	t.wg.Add(1)

	go func() {
		var err error

		t.concurrency <- struct{}{}

		defer func() {
			if err != nil {
				t.lock.Lock()
				t.err = multierr.Append(t.err, err)
				t.lock.Unlock()
			}
			<-t.concurrency
			t.wg.Done()
		}()


		select {
		case <-t.ctx.Done():
			// abort the loop if the context is dead
			t.cancelledOnce.Do(func() {
				err = t.ctx.Err()
			})
			return
		default:
		}

		err = f()
	}()

}

// Wait blocks until all function calls from the Go method have returned.
func (t *ThrottledExecutor) Wait() error {
	t.wg.Wait()
	return t.err
}
