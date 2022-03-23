package main

import "fmt"

/*功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）


数据范围： 1 \le n \le 2 \times 10^{9} + 14 \1≤n≤2×10
9
+14
输入描述：
输入一个整数

输出描述：
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。*/

func main() {
	var num int
	for {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			return
		}

		if num < 2 {
			fmt.Println(num)
			continue
		}

		for i := 2; i*i <= num; {
			if num%i == 0 {
				num /= i
				fmt.Print(i, " ")
			} else {
				i++
			}
		}
		if num > 1 {
			fmt.Println(num, " ")
		}
	}
}
