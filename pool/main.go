package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

// 线程池 https://github.com/panjf2000/ants

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFun() {
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Hello World !")
}
func main() {
	defer ants.Release()
	runTimes := 1000

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFun()
		wg.Done()
	}
	for run := 0; run < runTimes; run++ {
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Println("finished all tasks.")

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finished all tasks, result is %d\n", sum)

}
