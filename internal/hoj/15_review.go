package main

import "fmt"

/*描述
输入一个 int 型的正整数，计算出该 int 型数据在内存中存储时 1 的个数。

数据范围：保证在 32 位整型数字范围内
输入描述：
输入一个整数（int类型）

输出描述：
这个数转换成2进制后，输出1的个数*/

func main() {
	var num int
	_, err := fmt.Scanf("%d\n", &num)
	if err != nil {
		return
	}
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	fmt.Println(count)
}
