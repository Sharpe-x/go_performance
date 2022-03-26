package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*密码要求:

1.长度超过8位

2.包括大小写字母.数字.其它符号,以上四种至少三种

3.不能有长度大于2的不含公共元素的子串重复 （注：其他符号不含空格或换行）

数据范围：输入的字符串长度满足 1 \le n \le 100 \1≤n≤100
输入描述：
一组字符串。

输出描述：
如果符合要求输出：OK，否则输出NG*/
const (
	ok = "OK"
	ng = "NG"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		str := scanner.Text()
		if len(str) <= 8 {
			fmt.Println(ng)
			continue
		}

		var upWord, lowWord, num, special int
		for _, s := range str {
			switch {
			case s >= 'a' && s <= 'z':
				lowWord = 1
			case s >= 'A' && s <= 'Z':
				upWord = 1
			case s >= '0' && s <= '9':
				num = 1
			default:
				special = 1
			}
		}

		if upWord+lowWord+num+special < 3 {
			fmt.Println(ng)
			continue
		}

		isValid := true
		for i := 0; i < len(str)-3; i++ {
			if strings.Count(str, str[i:i+3]) > 1 {
				isValid = false
				break
			}
		}

		if isValid {
			fmt.Println(ok)
		} else {
			fmt.Println(ng)
		}

	}
}
