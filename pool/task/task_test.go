package task

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	myWorker()
}
func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, errors chan<- error) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		time.Sleep(time.Second)

		if job%2 == 0 {
			results <- job * 2
		} else {
			errors <- fmt.Errorf("error on job %v", job)
		}
		wg.Done()
	}
}

func myWorker() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	errors := make(chan error, 100)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		go worker(i, &wg, jobs, results, errors)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()

	select {
	case err := <-errors:
		fmt.Println("finished with error:", err.Error())
	default:

	}
}
