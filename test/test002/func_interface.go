package main

import "fmt"

type Printer interface {
	print(interface{})
}

//函数定义为类型

type FuncCaller func(p interface{})

// 实现接口的print()方法

func (fun FuncCaller) print(p interface{}) {
	// 调用FuncCaller函数体
	fun(p)
}
func main() {
	var printer Printer
	// 将匿名函数强转为FuncCaller 赋值给Printer
	printer = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	printer.print("Helle world!")
}
