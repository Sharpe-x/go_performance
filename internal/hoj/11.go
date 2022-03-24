package main

import "fmt"

/*描述
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001


数据范围： 0 \le n \le 2^{30}-1 \0≤n≤2
30
−1
输入描述：
输入一个int整数

输出描述：
将这个整数以字符串的形式逆序输出*/

func main() {
	var num int
	_, err := fmt.Scanf("%d\n", &num)
	if err != nil {
		return
	}
	if num < 0 {
		return
	}
	if num == 0 {
		fmt.Println(0)
	}

	for num > 0 {
		m := num % 10
		num = num / 10
		fmt.Print(m)
	}
	fmt.Println()
}
