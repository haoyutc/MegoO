package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// 1、每秒钟调用一次proc()
		// 2、程序不能退出
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}
