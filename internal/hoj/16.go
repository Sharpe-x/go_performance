package main

import "fmt"

func main() {
	var N, m int
	_, err := fmt.Scanf("%d %d\n", &N, &m)
	if err != nil {
		return
	}

	var v, w [61][3]int
	var price, weight, mainIndex int
	var memo [61][32000]int

	for i := 1; i <= m; i++ {
		_, err = fmt.Scanf("%d %d %d\n", &price, &weight, &mainIndex)
		if err != nil {
			return
		}
		if mainIndex != 0 {
			if w[mainIndex][1] != 0 { // 附件2
				v[mainIndex][2] = price
				w[mainIndex][2] = price * weight
			} else { // 附件一
				v[mainIndex][1] = price
				w[mainIndex][1] = price * weight
			}
		} else { // 主件
			v[i][0] = price
			w[i][0] = price * weight
		}

	}
	fmt.Println(getMax(&v, &w, m, N, &memo))
}

func getMax(v, w *[61][3]int, m int, N int, memo *[61][32000]int) int {

	if N <= 0 || m < 0 {
		return 0
	}

	// 记忆话搜索
	if memo[m][N] != 0 {
		return memo[m][N]
	}

	// m 不放
	res := getMax(v, w, m-1, N, memo)

	// m放主件
	if N >= v[m][0] {
		res = max(res, getMax(v, w, m-1, N-v[m][0], memo)+w[m][0])
	}

	// 主件 + 附件1
	price := v[m][0] + v[m][1]
	if N >= price {
		res = max(res, getMax(v, w, m-1, N-price, memo)+w[m][0]+w[m][1])
	}
	//  主件 + 附件2
	price = v[m][0] + v[m][2]
	if N >= price {
		res = max(res, getMax(v, w, m-1, N-price, memo)+w[m][0]+w[m][2])
	}

	// 主件 + 附件1 + 附件2
	price = v[m][0] + v[m][1] + v[m][2]
	if N >= price {
		res = max(res, getMax(v, w, m-1, N-price, memo)+w[m][0]+w[m][1]+w[m][2])
	}
	memo[m][N] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
