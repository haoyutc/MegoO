package main

import "fmt"

/*
*
golang 中的切片底层其实使用的是数组。当使用 str1[1:]时， str2 和 str1 底层共 享一个数组，这会导致 str2[1] = "new" 语句影响 str1 。
而 append 会导致底层数组扩容，生成新的数组，因此追加数据后的 str2 不会影 响 str1 。
但是为什么对 str2 复制后影响的却是 str1 的第三个元素呢?这是因为切
片 str2 是从数组的第二个元素开始， str2 索引为 1 的元素对应的是 str1 索引为 2 的元素。
result:
[a b c]
[b c]
[a b new]
[b new]
[a b new]
[b new z x y]
*/
func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	fmt.Println(str1) // [a b c]
	fmt.Println(str2) // [b c]
	str2[1] = "new"
	fmt.Println(str1) // [a b new]
	fmt.Println(str2) // [b new]
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1) // [a b new]
	fmt.Println(str2) // [b new z x y]
}
