package main

import (
	"fmt"
	"runtime"
	"sync"
)

var AppVersion string

func main() {
	//appVersion()
	//testByte()
	//defer_call()
	//goroutines()
	//calculate()
	slice()
}
func appVersion() {
	fmt.Println("Hello world. Version: ", AppVersion)
	fmt.Println(`Version: ` + AppVersion)
}
func testByte() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {

		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func goroutines() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i = ", i)
		}()
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i = ", i)
		}(i)
		wg.Done()
	}
	wg.Wait()
}

func calculate() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func slice() {
	sli := make([]int, 5)
	sli = append(sli, 1, 2, 3)
	fmt.Println(sli)
}
