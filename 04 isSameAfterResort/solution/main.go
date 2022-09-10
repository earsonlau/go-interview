package main

import (
	"fmt"
	"strings"
)

// 问题1描述

// 给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 这里规定【大小写为不同字符】，且考虑字符串中空格。
// 给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

// 思路
// 首先要保证字符串长度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。
func isResort(s1, s2 string) bool {
	length_s1 := len([]rune(s1))
	length_s2 := len([]rune(s2))

	if length_s1 > 5000 || length_s2 > 5000 || length_s1 != length_s2 {
		return false
	}

	for _,v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	// 使用strings.Count()
	fmt.Println(isResort("a123","321a"))
}
