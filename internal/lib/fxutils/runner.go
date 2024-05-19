package fxutils

import (
	"sync"
	"sync/atomic"
)

const (
	runnerIsWaitingFalse = int32(0)
	runnerIsWaitingTrue  = int32(1)
)

type Runner struct {
	wg        *sync.WaitGroup
	trackers  *sync.Map
	waitCH    chan struct{}
	isWaiting atomic.Int32
}

func NewRunner() *Runner {
	runner := &Runner{
		wg:       &sync.WaitGroup{},
		waitCH:   make(chan struct{}, 1),
		trackers: new(sync.Map),
	}

	runner.isWaiting.Store(runnerIsWaitingFalse)

	return runner
}

func (r *Runner) StartTracking(name string) {
	_, exist := r.trackers.LoadOrStore(name, struct{}{})
	if exist {
		return
	}

	r.wg.Add(1)

	if r.isWaiting.CompareAndSwap(runnerIsWaitingFalse, runnerIsWaitingTrue) {
		go r.waitTracking()
	}
}

func (r *Runner) StopTracking(name string) {
	r.trackers.CompareAndDelete(name, struct{}{})
	if val := r.isWaiting.Load(); val == runnerIsWaitingFalse {
		return
	}

	r.wg.Done()
}

func (r *Runner) waitChannel() <-chan struct{} {
	return r.waitCH
}

func (r *Runner) waitTracking() {
	r.wg.Wait()
	r.waitCH <- struct{}{}

	r.isWaiting.Swap(runnerIsWaitingFalse)
}
