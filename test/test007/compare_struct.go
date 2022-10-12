package main

import "fmt"

type Teacher struct {
	Name string
}

func main() {
	fmt.Println(&Teacher{Name: "wangwu"} == &Teacher{Name: "wangwu"})
	fmt.Println(Teacher{Name: "wangwu"} == Teacher{Name: "wangwu"})

	fmt.Println([...]string{"1"} == [...]string{"1"})
	//fmt.Println([]string{"1"}==[]string{"1"})
}
