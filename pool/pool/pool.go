package pool

import (
	"github.com/megoo/pool/task"
	"sync"
)

type Pool struct {
	Tasks       []*task.Task
	concurrency int
	tackChan    chan *task.Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*task.Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tackChan:    make(chan *task.Task),
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tackChan <- task
	}
	// all workers return
	close(p.tackChan)
	p.wg.Wait()
}

func (p *Pool) work() {
	for task := range p.tackChan {
		task.Run(&p.wg)
	}
}
