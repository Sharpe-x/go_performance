package main

import "strings"

/*问题描述

给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个 string s1 和一个 string s2，请返回一个 bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于 5000。

解题思路

首先要保证字符串长度小于 5000。之后只需要一次循环遍历 s1 中的字符在 s2 是否都存在即可*/

func isRegroup(s1, s2 string) bool {
	str1 := []rune(s1)
	str2 := []rune(s2)
	if len(str1) != len(str2) {
		return false
	}

	for _, s := range str1 {
		if strings.Count(s1, string(s)) != strings.Count(s2, string(s)) {
			return false
		}
	}

	return true
}
