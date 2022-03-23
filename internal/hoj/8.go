package main

import (
	"fmt"
	"sort"
)

/*描述
数据表记录包含表索引index和数值value（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照index值升序进行输出。


提示:
0 <= index <= 11111111
1 <= value <= 100000

输入描述：
先输入键值对的个数n（1 <= n <= 500）
接下来n行每行输入成对的index和value值，以空格隔开

输出描述：
输出合并后的键值对（多行）
*/

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	for {
		var nums int
		_, err := fmt.Scanf("%d", &nums)
		if err != nil {
			break
		}
		var index, value int
		m := make(map[int]int)
		for i := 0; i < nums; i++ {
			_, err = fmt.Scanf("%d %d\n", &index, &value)
			if err != nil {
				break
			}
			if m[index] == 0 {
				m[index] = value
			} else {
				m[index] += value
			}
		}

		var res []int
		for key := range m {
			res = append(res, key)
		}
		sort.Sort(sort.IntSlice(res))
		for _, k := range res {
			fmt.Println(k, m[k])
		}
	}
}
