package main

import "fmt"

// 问题1描述

// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

// 思路
// 翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。

func reverseString(s string) (string,bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i ++ {
		str[i],str[l-1-i] = str[l-1-i],str[i]
	}
	return string(str),true
}

func main() {
	// 使用strings.Count()
	fmt.Println(reverseString("aaaa"))
}
