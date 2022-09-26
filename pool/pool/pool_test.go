package pool

import (
	"github.com/megoo/pool/task"
	"github.com/micro/go-micro/debug/log"
	"testing"
)

func TestPool(t *testing.T) {
	myPool()
}

func myPool() {
	tasks := []*task.Task{
		task.NewTask(func() error {
			return nil
		}),
		task.NewTask(func() error {
			return nil
		}),
		task.NewTask(func() error {
			return nil
		}),
	}
	p := NewPool(tasks, 3)
	p.Run()
	var numErrs int
	for _, task := range tasks {
		if task.Err != nil {
			log.Error(task.Err)
			numErrs++
		}
		if numErrs >= 10 {
			log.Error("Too many errors.")
			break
		}
	}
}
