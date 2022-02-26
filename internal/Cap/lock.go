package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	count int32
	mutex sync.Mutex
)

type Lock struct {
	c chan struct{}
}

func NewLock() *Lock {
	channel := make(chan struct{}, 1)
	channel <- struct{}{}
	return &Lock{
		c: channel,
	}
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
		lockResult = false
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

func noLock() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func lock() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func tryLock() {
	lock := NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				println("lock failed")
				return
			}
			count++
			println("current count= ", count)
			lock.Unlock()
		}()
	}
	wg.Wait()

}

func cas() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func main() {
	noLock()
	count = 0
	lock()
	count = 0
	tryLock()
	count = 0
	cas()
}
