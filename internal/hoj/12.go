package main

import (
	"fmt"
)

/*
描述
接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述：
输入一行，为一个只包含小写字母的字符串。

输出描述：
输出该字符串反转后的字符串。*/

func main() {
	var s string
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		return
	}
	bytes := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		bytes[i], bytes[len(s)-i-1] = bytes[len(s)-i-1], bytes[i]
	}
	fmt.Println(string(bytes))
}
