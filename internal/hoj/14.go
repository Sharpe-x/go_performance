package main

import (
	"fmt"
	"sort"
)

/*描述
给定 n 个字符串，请对 n 个字符串按照字典序排列。

数据范围： 1 \le n \le 1000 \1≤n≤1000  ，字符串长度满足 1 \le len \le 100 \1≤len≤100
输入描述：
输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
输出描述：
数据输出n行，输出结果为按照字典序排列的字符串。*/
func main() {
	var num int
	_, err := fmt.Scanf("%d\n", &num)
	if err != nil {
		return
	}

	var (
		s    string
		strs []string
	)
	for i := 0; i < num; i++ {
		_, err = fmt.Scanf("%s\n", &s)
		if err != nil {
			return
		}
		strs = append(strs, s)
	}
	/*	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})*/

	sort.Strings(strs)
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
}
