package test

import (
	"fmt"
	"reflect"
	"testing"
)

// 函数返回值
//在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。 如果返回值有有多个返回值必须加上括号；
//如果只有一个返回值并且有命名也需要加上括号； 此处函数第一个返回值有sum名称，第二个未命名，所以错误。

func TestReturnVal(t *testing.T) {
	num, _ := myFun(1, 2)
	fmt.Println("sum = ", num)
}

func myFun(x, y int) (sum int, err error) {
	return x + y, nil
}

// 结构体比较

func TestStructCompare(t *testing.T) {
	p1 := struct {
		age  int
		name string
	}{age: 23, name: "lisi"}
	p2 := struct {
		age  int
		name string
	}{age: 23, name: "lisi"}

	if p1 == p2 {
		fmt.Println("p1==p2")
	}

	/*m1 := struct {
		age   int
		name  string
		hobby map[string]string
	}{age: 24, name: "wangwu", hobby: map[string]string{"one": "swimming"}}
	m2 := struct {
		age   int
		name  string
		hobby map[string]string
	}{age: 24, name: "wangwu", hobby: map[string]string{"one": "swimming"}}

	if m1 == m2 { // 编译不通过: invalid operation: m1 == m2 (struct containing map[string]string cannot be compared)
		fmt.Println("m1==m2")
	}*/

	// 结构体比较规则注意1：只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关.
	/*w1 := struct {
		age  int
		name string
	}{age: 22, name: "lisi"}
	w2 := struct {
		name string
		age  int
	}{name: "lisi", age: 22}
	if w1 == w2 {// 编译不通过：字段顺序不同，不是相同的结构体了，不能比较
		fmt.Println("w1==w2")
	}*/

	// 结构体比较规则注意2：结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice，则结构体不能用==比较。
	// 可以使用reflect.DeepEqual进行比较
	m1 := struct {
		age   int
		name  string
		hobby map[string]string
	}{age: 24, name: "wangwu", hobby: map[string]string{"one": "swimming"}}
	m2 := struct {
		age   int
		name  string
		hobby map[string]string
	}{age: 24, name: "wangwu", hobby: map[string]string{"one": "swimming"}}

	if reflect.DeepEqual(m1, m2) {
		fmt.Println("m1==m2")
	} else {
		fmt.Println("m1 != m2")
	}
}

// string和nil
// nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。
// 通常编译的时候不会报错，但是运行是时候会报:cannot use nil as type string in return argument

func TestStringNil(t *testing.T) {
	m := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, ok := ISExist(m, 3)
	fmt.Println(v, ok)
}

func ISExist(m map[int]string, id int) (string, bool) {
	if _, ok := m[id]; !ok {
		//return nil, false // 编译不会通过: Cannot use 'nil' as the type string
		return "value not exist", false
	}
	return "value exist", true
}

// 切片追加, make初始化均为0
// new和make的区别：
//
// 二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；
//而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。
//
// make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
func TestSliceAppend(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4]

	list := new([]int)
	*list = append(*list, 1)
	fmt.Println(*list)
}

type student struct {
	Id   int
	Name string
	Age  int
}

func TestMapForeach(t *testing.T) {
	stus := []student{{
		Id: 1, Name: "lisi", Age: 22,
	}, {
		Id: 2, Name: "wangwu", Age: 22,
	}, {
		Id: 3, Name: "wangba", Age: 22,
	},
	}
	stusmap := make(map[int]student)

	for _, s := range stus {
		stusmap[s.Id] = s
	}

	for k, v := range stusmap {
		fmt.Println(k, "=>", v.Name)
	}

}
func TestMapForeach1(t *testing.T) {
	stus := []student{{
		Id: 1, Name: "lisi", Age: 22,
	}, {
		Id: 2, Name: "wangwu", Age: 22,
	}, {
		Id: 3, Name: "wangba", Age: 22,
	},
	}
	stusmap := make(map[int]*student)

	for _, s := range stus {
		stusmap[s.Id] = &s
	}

	for k, v := range stusmap {
		fmt.Println(k, "=>", v.Name)
	}

}
func TestMapForeach2(t *testing.T) {
	stus := []student{{
		Id: 1, Name: "lisi", Age: 22,
	}, {
		Id: 2, Name: "wangwu", Age: 22,
	}, {
		Id: 3, Name: "wangba", Age: 22,
	},
	}
	stusmap := make(map[int]*student)

	for i := 0; i < len(stus); i++ {
		stusmap[stus[i].Id] = &stus[i]
	}

	for k, v := range stusmap {
		fmt.Println(k, "=>", v.Name)
	}

}

type People interface {
	Speak(string) string
}

type Worker struct {
}

func (w *Worker) Speak(s string) (talk string) {
	if s == "love" {
		talk = "you are a good worker"
	} else {
		talk = "Hi"
	}
	return
}

func TestSpeak(t *testing.T) {
	var peo People = &Worker{}
	s := "love"
	fmt.Println(peo.Speak(s))
}
