package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 问题1描述

// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值
// true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

// 思路
// 这里有几个重点，第一个是ASCII字符，ASCII字符字符一共有256个
// 其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。
// 然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。
// 如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。

// 使用strings.Count()
func isUniqueString_strings_count(s string) bool{
	if strings.Count(s, "") > 3000{
		return false
	}
	for _,v := range s{
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1{
			return false
		}
	}
	return true
}


// 使用strings.Index()
func isUniqueString_strings_index(s string) bool{
	if strings.Count(s, "") > 3000{
		return false
	}
	for k,v := range s{
		if v > 127 {
			return false
		}
		if strings.Index(s,string(v)) != k{
			return false
		}
	}
	return true
}

// 使用bitmap
func isUniqueString_bit_arithmetic(s string) bool{
	if len(s) == 0 || len(s) > 3000 {
		return false
	}
	// 256个字符= 4 * 64
	var m1,m2,m3,m4 uint64
	var mark *uint64
	for _, r := range s {
		// r先强制转换为长度为64的数字
		n := uint64(r)
		if n < 64{
			mark = &m1
		}else if n <128 {
			mark = &m2
			// n映射回[0,63]
			n -= 64
		}else if n < 192{
			mark = &m3
			// n映射回[0,63]
			n -= 128
		}else {
			mark = &m4
			// n映射回[0,63]
			n -= 192
		}
		// 判断mark指针指向的变量里面的位跟(1<<n)有无重复，有重复说明对应字符重复
		if (*mark) & (1 << n) != 0 {
			return false
		}
		// 或操作，把1加到mark指针指向的变量对应的位置里面
		*mark = (*mark) | uint64(1 << n)
	}
	return true
}


func main(){
	fmt.Println(strconv.FormatBool(isUniqueString_strings_count("123"))) 
	fmt.Println(strconv.FormatBool(isUniqueString_strings_index("123"))) 
	fmt.Println(strconv.FormatBool(isUniqueString_bit_arithmetic("123"))) 
}