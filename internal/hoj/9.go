package main

import "fmt"

/*描述
输入一个 int 型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是 0 。

数据范围： 1 \le n \le 10^{8} \1≤n≤10
8

输入描述：
输入一个int型整数

输出描述：
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数*/

func main() {
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return
	}

	m := make(map[int]bool)
	res := 0
	for num > 0 {
		temp := num % 10
		if m[temp] == false {
			res = res*10 + temp
			m[temp] = true
		}
		num = num / 10
	}
	fmt.Println(res)
}
