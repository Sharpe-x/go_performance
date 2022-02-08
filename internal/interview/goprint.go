package main

import (
	"fmt"
	"sync"
)

/*
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

*/

func main() {
	letter, number := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- struct{}{}
				break
			default:
				break
			}
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= len(str) {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				number <- struct{}{}
				break
			default:
				break
			}
		}
	}(&wg)
	number <- struct{}{}
	wg.Wait()
}
