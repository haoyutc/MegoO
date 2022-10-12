package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, ch)
	}
	if TimeWait(&wg, time.Second*5) {
		close(ch)
		fmt.Println("Timeout exit!")
	}
	time.Sleep(time.Second * 10)
}

// 要求sync.WaitGroup支持timeout功能
//如果timeout到了超时时间返回true
//如果WaitGroup自然结束返回false

func TimeWait(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() {
		ch <- true
	})
	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
