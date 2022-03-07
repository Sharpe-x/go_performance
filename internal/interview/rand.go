package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(10000)
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
