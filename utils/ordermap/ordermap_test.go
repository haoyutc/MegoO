package ordermap

import (
	"fmt"
	"github.com/elliotchance/orderedmap/v2"
	"testing"
)

func TestMapOrderForeach(t *testing.T) {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func TestOrderMap(t *testing.T) {
	m := orderedmap.NewOrderedMap[string, any]()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set("123", true)

	for _, k := range m.Keys() {
		value, _ := m.Get(k)
		fmt.Println(k, value)
	}

	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
	m.Delete("qux")
}
