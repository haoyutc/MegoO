package job

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	go func() {
		time.Sleep(3 * time.Second)
		println("World") // <------ this will never execute because the program will have already exited!
	}()
	println("Hello")
}

func TestRoutine2(t *testing.T) {
	finished := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		println("World")
		finished <- true
	}()
	println("Hello")
	<-finished
}

func TestRoutine3(t *testing.T) {
	worldChannel := make(chan string)
	dearChannel := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		worldChannel <- "world"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		dearChannel <- "dear"
	}()
	println("Hello", <-dearChannel, <-worldChannel)
}

func TestJob(t *testing.T) {
	jobs := make([]*Job, 0)
	for i := 0; i < 10; i++ {
		jobs = append(jobs, &Job{Id: i})
	}
	jobRets := launchAndReturn(jobs)
	fmt.Println(jobRets)
}

func TestRunner(t *testing.T) {
	runner := NewRunner()
	defer runner.Stop()
	time.Sleep(3 * time.Second)
	runner.Run()
}

func TestJobLaunchThrottleAndForget(t *testing.T) {
	start := time.Now()
	var jobs []Job
	for i := 0; i < 10; i++ {
		jobs = append(jobs, Job{Id: i})
	}

	var (
		wg         sync.WaitGroup
		jobChannel chan Job
	)

	wg.Add(NumberOfWorkers)
	jobChannel = make(chan Job, 0)

	// start the workers
	for i := 0; i < NumberOfWorkers; i++ {
		go worker(i, &wg, jobChannel)
	}

	// send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	fmt.Printf("Took %s\n", time.Since(start))

}

func TestJobLaunchThrottleAndReturn(t *testing.T) {
	start := time.Now()
	var jobs []Job
	for i := 0; i < 10; i++ {
		jobs = append(jobs, Job{Id: i})
	}

	var (
		wg               sync.WaitGroup
		jobChannel       chan Job
		jobResultChannel chan JobResult
		jobResults       []JobResult
	)

	jobChannel = make(chan Job)
	jobResultChannel = make(chan JobResult, len(jobs))
	jobResults = make([]JobResult, 0)
	wg.Add(NumberOfWorkers)

	// start the workers
	for i := 0; i < NumberOfWorkers; i++ {
		go worker2(i, &wg, jobChannel, jobResultChannel)
	}

	// send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	close(jobResultChannel)

	// receive job results from workers
	for result := range jobResultChannel {
		jobResults = append(jobResults, result)
	}
	fmt.Println(jobResults)
	fmt.Printf("Took %s\n", time.Since(start))

}

func TestTimeBaseThrottling(t *testing.T) {
	start := time.Now()
	// crate fake jobs for testing purposes
	var jobs []Job
	for i := 0; i < 10; i++ {
		jobs = append(jobs, Job{Id: i})
	}
	var (
		wg         *sync.WaitGroup
		jobChannel = make(chan Job)
	)
	wg = &sync.WaitGroup{}
	wg.Add(NumberOfWorkers)

	// start the workers
	for i := 0; i < NumberOfWorkers; i++ {
		go worker3(i, wg, jobChannel)
	}
	// send job to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	fmt.Printf("Took %s\n", time.Since(start))
}

func TestUsingChannelToPerformBulkOperations(t *testing.T) {
	jobChannel := make(chan Job, 1000)
	go worker4(jobChannel)
	for i := 0; i < 5000; i++ {
		jobChannel <- Job{Id: i}
	}
	// wait for channel to be empty
	for len(jobChannel) != 0 {
		time.Sleep(100 * time.Millisecond)
	}
}
