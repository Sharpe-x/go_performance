package main

import (
	"fmt"
	"sort"
)

// 从高位开始遍历，对每一位先尝试使用相同数字，除了最后一位。
// 如果没有相同的数字时，尝试是否有比当前数字更小的，有的话选更小的数字里最大的，剩下的用最大数字。
//都没有就向前回溯看前一个有没有更小的。如果一直回溯到第一个数字都没有更小的数字，就用位数更少的全都是最大数字的数。

// 获取小于指定数字的最大数字
// 数组有序 从小到大
// https://blog.csdn.net/K346K346/article/details/126958310
func getMaxDigitLtD(digits []int, digint int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digint {
			return digits[i]
		}
	}
	return 0
}

func getMaxNumLtN(digits []int, n int) int {
	var ndigits []int
	// 获取n的每一位数字
	for n > 0 {
		ndigits = append(ndigits, n%10)
		n /= 10
	}

	sort.Ints(digits) // 排序给定数组

	// 数字写入map 用来判断是否存在
	m := make(map[int]struct{})
	for _, v := range digits {
		m[v] = struct{}{}
	}

	// 目标数的每一位数字
	tdigits := make([]int, len(ndigits))

	// 从高位遍历，尽可能取相同值 除了最后一位
	for i := len(ndigits) - 1; i >= 0; i-- {
		if _, ok := m[ndigits[i]]; ok && i > 0 {
			tdigits[i] = ndigits[i]
			continue
		}

		// 存在小于当前位的最大数字
		if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
			tdigits[i] = d
			break
		}

		//回溯
		for i++; i < len(ndigits); i++ {
			tdigits[i] = 0
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				tdigits[i] = d
				break
			}
			if i == len(ndigits)-1 {
				tdigits = tdigits[:len(tdigits)-1]
			}
		}
		break
	}

	var target int
	for i := len(tdigits) - 1; i >= 0; i-- {
		target *= 10
		if tdigits[i] > 0 {
			target += tdigits[i]
			continue
		}
		target += digits[len(digits)-1]
	}
	return target
}

func main() {
	fmt.Println(getMaxNumLtN([]int{1, 2, 9, 4}, 2533))
	fmt.Println(getMaxNumLtN([]int{1, 2, 5, 4}, 2543))
	fmt.Println(getMaxNumLtN([]int{1, 2, 5, 4}, 2541))
	fmt.Println(getMaxNumLtN([]int{1, 2, 9, 4}, 2111))
	fmt.Println(getMaxNumLtN([]int{5, 9}, 5555))
	fmt.Println(getMaxNumLtN([]int{1, 8}, 255))
	fmt.Println(getMaxNumLtN([]int{3, 8}, 255))
}
