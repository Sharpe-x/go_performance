package concurrency

import (
	"sync"
	"time"
)

// Go 语言标准库 sync 提供了 2 种锁，互斥锁(sync.Mutex)和读写锁(sync.RWMutex)
// 互斥即不可同时运行。即使用了互斥锁的两个代码片段互相排斥，只有其中一个代码片段执行完成后，另一个才可以执行
// sync 标准库中提供了 sync.Mutex  互斥锁类型及其两个方法；
// Lock 加锁
// Unlock 释放锁
// 我们可以通过在代码前调用Lock 方法 在代码后调用Unlock 方法来保证一段代码的互斥执行 也可以用defer 语句来保证互斥锁一定会被解锁。
// 在一个Go 协程调用Lock方法获得锁后，其他请求锁的协程都会阻塞在Lock 方法。直到锁被释放。

// 读写锁 sync.RWMutex 读锁是允许同时执行的 但写锁是互斥的
// 读锁之间不互斥 没有写锁的情况下 读锁是无阻塞的 多个协程可以同时获得读锁
// 写锁之间是互斥的 存在写锁 其他写锁阻塞
// 写锁和读锁是互斥的 如果存在读锁 写锁阻塞 如果存在写锁 读锁阻塞

// sync.RWMutex 互斥锁类型及其四个方法
// Lock 加锁
// Unlock 释放写锁
// RLock 加读锁
// RUnlock 释放读锁
// 读写锁的存在是为了解决读多写少是的性能问题 读场景较多时 读写锁可有效地减少锁阻塞的时间。

// 互斥锁如何实现公平
// https://colobu.com/2018/12/18/dive-into-sync-mutex/

// 互斥锁有两种状态：正常状态和饥饿状态。
// 在正常状态下，所有等待锁的 goroutine 按照FIFO顺序等待。唤醒的 goroutine 不会直接拥有锁，
//而是会和新请求锁的 goroutine 竞争锁的拥有。新请求锁的 goroutine 具有优势：它正在 CPU 上执行，而且可能有好几个，
//所以刚刚唤醒的 goroutine 有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的 goroutine 会加入到等待队列的前面。
//如果一个等待的 goroutine 超过 1ms 没有获取锁，那么它将会把锁转变为饥饿模式。

// 在饥饿模式下，锁的所有权将从 unlock 的 goroutine 直接交给交给等待队列中的第一个。
//新来的 goroutine 将不会尝试去获得锁，即使锁看起来是 unlock 状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。

// 如果一个等待的 goroutine 获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个；(2)它等待的时候小于1ms。它会将锁的状态转换为正常状态。
// 正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (r *RWLock) Write() {
	r.mu.Lock()
	r.count++
	time.Sleep(cost)
	r.mu.Unlock()
}

func (r *RWLock) Read() {
	r.mu.RLock()
	_ = r.count
	time.Sleep(cost)
	r.mu.RUnlock()
}
