package main

import (
	"fmt"
	"sync"
	"time"
)

// GO 里面 MAP 如何实现 key 不存在 get 操作等待 直到 key 存在或者超时，保证并发安全，且需要实现以下接口

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

// 并发安全，那么必须用锁 多个 goroutine 读的时候如果值不存在则阻塞，直到写入值，那么每个键值需要有一个阻塞 goroutine 的 channel。

func newMap() *Map {
	return &Map{
		c: make(map[string]*entry),
	}
}

type Map struct {
	c   map[string]*entry
	rmx sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		//e.ch <- struct{}{}
		close(e.ch)
	} else {
		e = &entry{ch: make(chan struct{}), isExist: true, value: val}
		m.c[key] = e
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.Unlock()
		return e.value
	} else if !ok {
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		fmt.Println("!ok 协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("!ok 协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	testMap := newMap()
	testMap.Out("hello", 1)
	testMap.Out("world", 2)

	go func() {
		fmt.Println(testMap.Rd("hi", time.Second*10))
	}()
	// go fmt.Println(testMap.Rd("hi", time.Second*10))
	// 上下2种有区别 go fmt.Println(testMap.Rd("hi", time.Second*10)) 会把内部参数算出来

	time.Sleep(time.Second * 3)
	testMap.Out("hi", 100)
	time.Sleep(time.Second * 3)
}
