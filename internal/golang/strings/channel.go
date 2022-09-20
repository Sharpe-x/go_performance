package main

import (
	"fmt"
	"sync"
)

func main() {

	ansChan := make(chan int, 10)
	var ans []int

	go func() {
		for num := range ansChan {
			ans = append(ans, num)
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		num := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			ansChan <- num
		}()
	}
	wg.Wait()
	close(ansChan)
	fmt.Println(ans)

}
