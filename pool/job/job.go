package job

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// https://twin.sh/articles/39/go-concurrency-goroutines-worker-pools-and-throttling-made-simple

type Job struct {
	Id int
}

type JobResult struct {
	Output string
}

func launchAndReturn(jobs []*Job) []JobResult {
	var (
		results      []JobResult
		resultsMutex sync.RWMutex
		wg           sync.WaitGroup
	)

	wg.Add(len(jobs))
	for _, job := range jobs {
		go func(job *Job) {
			defer wg.Done()
			jobRet := doMyBusiness(job)
			resultsMutex.Lock()
			results = append(results, jobRet)
			resultsMutex.Unlock()
		}(job)
	}
	wg.Wait()
	return results
}

func doMyBusiness(job *Job) JobResult {
	fmt.Printf("Running job #%d\n", job.Id)
	return JobResult{Output: "Success"}
}

func doSomeThing(workerId int, job Job) JobResult {
	fmt.Printf("Worker #%d Running job #%d\n", workerId, job.Id)
	time.Sleep(time.Second)
	return JobResult{Output: "Success"}
}

func doSomeThing2(workerId int, job Job) JobResult {
	fmt.Printf("Worker #%d Running job #%d\n", workerId, job.Id)
	time.Sleep(500 * time.Millisecond)
	return JobResult{Output: "Success"}
}

const NumberOfWorkers = 10

const MaximumNumberOfExecutionPerSecond = 50

func worker(id int, wg *sync.WaitGroup, jobChannel <-chan Job) {
	defer wg.Done()
	for job := range jobChannel {
		doSomeThing(id, job)
	}
}

func worker2(id int, wg *sync.WaitGroup, jobChannel <-chan Job, resultChannel chan JobResult) {
	defer wg.Done()
	for job := range jobChannel {
		resultChannel <- doSomeThing2(id, job)
	}
}

func worker3(id int, wg *sync.WaitGroup, jobChannel <-chan Job) {
	defer wg.Done()
	lastExecutionTime := time.Now()
	minimumTimeBetweenEachExecution := time.Duration(math.Ceil(1e9 / (MaximumNumberOfExecutionPerSecond / float64(NumberOfWorkers))))
	for job := range jobChannel {
		timeUtilNextExecution := -(time.Since(lastExecutionTime) - minimumTimeBetweenEachExecution)
		if timeUtilNextExecution > 0 {
			fmt.Printf("Worker #%d backing off for %s\n", id, timeUtilNextExecution.String())
			time.Sleep(timeUtilNextExecution)
		} else {
			fmt.Printf("Worker #%d not backing off\n", id)
		}
		lastExecutionTime = time.Now()
		doSomeThing3(id, job)
	}
}

func doSomeThing3(workerId int, job Job) JobResult {
	simulatedExecutionTime := rand.Intn(1000)
	fmt.Printf("Worker #%d Running job #%d (simulatedExecutionTime=%dms)\n", workerId, job.Id, simulatedExecutionTime)
	time.Sleep(time.Duration(simulatedExecutionTime) * time.Millisecond)
	return JobResult{Output: "Success"}
}

const MaxBulkSize = 50

func worker4(jobChannel <-chan Job) {
	var jobs []Job
	for true {
		if len(jobChannel) > 0 && len(jobs) < MaxBulkSize {
			jobs = append(jobs, <-jobChannel)
			continue
		}
		if (len(jobChannel) == 0 && len(jobs) == 0) || len(jobs) == MaxBulkSize {
			fmt.Printf("processing bulk of %d jobs\n", len(jobs))
			// clear the list of jobs that were just processed
			jobs = jobs[:0]
		}
		// no jobs in the channel ? back off
		if len(jobChannel) == 0 {
			fmt.Println("Backing off")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
