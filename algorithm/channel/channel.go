package main

import (
	"fmt"
	"time"
)

//func consume(ch chan int) {
//	// 线程休息100s 再从 channel 读取数据
//	time.Sleep(time.Second * 100)
//	<-ch
//}
//func main() {
//	// 创建一个长度为2的channel
//	ch := make(chan int, 2)
//	go consume(ch)
//
//	ch <- 0
//	ch <- 1
//
//	// 发送数据不被阻塞
//	fmt.Println("I'm free !")
//	ch <- 2
//	fmt.Println("I can not go there within 100s!")
//	time.Sleep(time.Second)
//}

func send(ch chan int, begin int) {
	// 循环向通道发送数据
	for i := begin; i < begin; i++ {
		ch <- i
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go send(ch2, 10)

	// 主 goroutine 休眠1秒，保证调度成功
	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1: // 从ch1	读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case val := <-ch2:
			fmt.Printf("get value %d from ch2\n", val)
		case <-time.After(2 * time.Second):
			fmt.Println("Time out")
			return
		}
	}
}
