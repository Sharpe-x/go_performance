package main

import "fmt"

/*问题描述

请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个 string，请返回一个 string，为翻转后的字符串。保证字符串的长度小于等于 5000。

解题思路

翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将 str[len]赋值给 str[0],将 str[0] 赋值 str[len]。*/

func reversString(str string) string {
	s := []rune(str)
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println("reversString(\"httpheeee\")", reversString("httpheeee"))
	fmt.Println("reversString(\"我来自中国China\")", reversString("我来自中国China"))
}
