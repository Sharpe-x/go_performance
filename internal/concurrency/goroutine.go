package concurrency

import (
	"fmt"
	"time"
)

// 超时控制在网络编程中是非常常见的，利用 context.WithTimeout 和 time.After 都能够很轻易地实现。

// time.After 实现超时控制
func doBadThings(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

// 利用 time.After 启动了一个异步的定时器，返回一个 channel，当超过指定的时间后，该 channel 将会接受到信号。
// 启动了子协程执行函数 f，函数执行结束后，将向 channel done 发送结束信号。
// 使用 select 阻塞等待 done 或 time.After 的信息，若超时，则返回错误，若没有超时，则返回 nil。

// 问题
// done 是一个无缓冲区的 channel，如果没有超时
//，doBadthing 中会向 done 发送信号，
//select 中会接收 done 的信号，因此 doBadthing 能够正常退出，子协程也能够正常退出。
// 当超时发生时，select 接收到 time.After 的超时信号就返回了，done 没有了接收方(receiver)，
//而 doBadthing 在执行 1s 后向 done 发送信号，由于没有接收者且无缓存区，发送者(sender)会一直阻塞，导致协程不能退出。
func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)

	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
	return nil
}

// 使用有缓冲chan避免发生阻塞
func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Microsecond):
		return fmt.Errorf("timeout")
	}
}

// 使用select 尝试发送
func doGoodThings(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}

// 因为 goroutine 不能被强制 kill，在超时或其他类似的场景下，为了 goroutine 尽可能正常退出，建议如下：
//
//尽量使用非阻塞 I/O（非阻塞 I/O 常用来实现高性能的网络库），阻塞 I/O 很可能导致 goroutine 在某个调用一直等待，而无法正确结束。
//业务逻辑总是考虑退出机制，避免死循环。
//任务分段执行，超时后即时退出，避免 goroutine 无用的执行过多，浪费资源。
