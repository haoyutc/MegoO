package job

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/debug/log"
	"sync"
	"time"
)

type Runner struct {
	ctx         context.Context
	cancelFunc  context.CancelFunc
	wg          *sync.WaitGroup
	errChannels []chan error
}

func NewRunner() *Runner {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &Runner{
		ctx:         ctx,
		cancelFunc:  cancelFunc,
		wg:          &sync.WaitGroup{},
		errChannels: make([]chan error, 0),
	}
}

func (r *Runner) Stop() {
	r.cancelFunc()
	r.wg.Wait()
}

func (r *Runner) Run() {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		timer := time.Tick(3 * time.Second)
		for {
			select {
			case <-timer:
				log.Info("Job is running")
				r.DoJob()
			case <-r.ctx.Done():
				return
			}
		}
	}()
}

func (r *Runner) DoJob() {
	fmt.Println("worker doing a task now")
}
