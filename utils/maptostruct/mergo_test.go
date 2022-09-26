package maptostruct

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
	"testing"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

// https://github.com/imdario/mergo

func TestMergo(t *testing.T) {
	var defaultStu = Student{
		Name: "zhangsan",
		Age:  22,
		Sex:  "man",
	}
	var m = make(map[string]interface{})
	if err := mergo.Map(&m, defaultStu); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("map m = %+v", m)
}
