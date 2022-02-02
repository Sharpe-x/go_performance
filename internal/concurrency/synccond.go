package concurrency

import (
	"log"
	"sync"
	"time"
)

// sync.Cond 条件变量用来协调想要访问共享资源的那些 goroutine 当共享资源的状态发生变化的时候 他可以用来通知被互斥锁阻塞的goroutine

/*sync.Cond 基于互斥锁/读写锁 他和互斥锁的区别是
互斥锁 sync.Mutex 通常用来保护临界区和共享资源 条件变量用来协调想要访问共享资源的goruntine

sync.Cond 经常用在多个goruntine 等待一个goruntine 通知的场景 如果是一个通知 一个等待 使用互斥锁或channel 就能搞定了*/

// 三个协程调用 Wait() 等待，另一个协程调用 Broadcast() 唤醒所有等待的协程。

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func condExample() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)

	write("write", cond)
	time.Sleep(time.Second * 3)
}
