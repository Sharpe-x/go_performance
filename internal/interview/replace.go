package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"unicode"
)

/*请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，
并且知道字符串的真实长度(小于等于 1000)，同时保证字符串由【大小写的英文字母组成】。 给定一个 string 为原始的串，返回替换后的 string。*/

func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}

func main() {
	fmt.Println(replaceBlank("1 2 3 4"))
	fmt.Println(replaceBlank("a bcdefg hi gk"))

	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
