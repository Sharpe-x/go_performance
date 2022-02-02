package concurrency

import (
	"log"
	"sync"
	"time"
)

// 利用信道 channel 的缓冲区大小来控制goroutine 数量
func useChannelControlGoroutineNum() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 30; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

// 利用第三方库
// https://github.com/Jeffail/tunny
// https://github.com/panjf2000/ants
