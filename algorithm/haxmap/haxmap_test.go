package haxmap

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"testing"
)

// https://github.com/alphadose/haxmap

func TestHaxMap(t *testing.T) {
	mep := haxmap.New[int, string]()
	mep.Set(1, "one")
	if val, ok := mep.Get(1); ok {
		println(val)
	}

	mep.Set(2, "two")
	mep.Set(3, "three")

	mep.ForEach(func(key int, value string) {
		fmt.Printf("key -> %d | value -> %s\n", key, value)
	})

	mep.Del(1)
	mep.Del(2)
	mep.Del(3)
	mep.Del(4)
	if mep.Len() == 0 {
		fmt.Println("cleanup complete!")
	}
}
