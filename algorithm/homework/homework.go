package homework

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// 交替打印数字和字母
/**
问题描述
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一 个 goroutine 打印字母， 最终效果如下:
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
解题思路
问题很简单，使用 channel 来控制打印的进度。
使用两个 channel ，来分别控制数字和 字母的打印序列， 数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数 字打印，然后周而复始的工作。
*/

func alternatePrintNumberAndString() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wg)
	number <- true
	wg.Wait()
}

// 判断字符串中字符是否全都不同
/**
问题描述
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允 许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都 不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的⻓ 度小于等于【3000】。
解题思路
这里有几个重点，第一个是 ASCII字符 ， ASCII字符 字符一共有256个，其中128个是常 用字符，可以在键盘上输入。
128之后的是键盘上无法找到的。 然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结 构，且字符串小于等于3000。
如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置 的方式实现。
*/
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}

// 翻转字符串
/**
问题描述
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字 符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的⻓度小于等于 5000。
解题思路
翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0], 将str[0] 赋值 str[len]。
*/

func reverseString(s string) (string, bool) {
	l := len(s)
	if l > 5000 {
		return s, false
	}
	chars := []rune(s)
	for i := 0; i < l/2; i++ {
		chars[i], chars[l-i-1] = chars[l-1-i], chars[i]
	}
	return string(chars), true
}

// 判断两个给定的字符串排序后是否一致
/**
问题描述
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一 个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
保证两串的⻓度都小于等于5000。
解题思路
首先要保证字符串⻓度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存 在即可。
*/

func isRegroup(s1, s2 string) bool {
	if len(s1) > 5000 || len(s2) > 5000 || len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

// 字符串替换问题
/**
问题描述
请编写一个方法，将字符串中的空格全部替换为“%20”。
假定该字符串有足够的空间存 放新增的字符，并且知道字符串的真实⻓度(小于等于1000)，同时保证字符串由【大小 写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。
解题思路
两个问题，第一个是只能是英文字母，第二个是替换空格。
*/

func replaceBlank(s string) (string, bool) {
	if len(s) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
