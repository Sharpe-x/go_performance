package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func do(taskCh chan int) {
	for { // 协程一直处于阻塞状态，等待接收任务
		select {
		case t := <-taskCh:
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTasks() {
	taskCh := make(chan int, 10)
	go do(taskCh)
	for i := 0; i < 1000; i++ { // 程序结束，协程也不会释放
		taskCh <- i
	}
}

// channel 的常见操作
/*ch := make(chan int) // 不带缓冲区
ch := make(chan int,10) // 带缓冲区  缓冲区满之前 即使没有接受方 发送方不阻塞
close(ch) // 关闭channel

ch <- v // 向通道发送v
<- ch // 忽略接受值
v := <- ch // 接受值并赋值给变量v
v,beforeClosed := <- ch //接收操作可以有2个返回值  beforeClosed 代表v 是不是信道关闭前发送的 false 表示信道已经关闭  true 代表是信道关闭前发送的
如果一个信道已经关闭 <- ch 将永远不会发生阻塞 但是我们可以通过第二个返回值beforeClosed 得知信道已经关闭 作出相应的处理*/

/*len(ch) 和cap(ch) 查询长度和容量*/

/*channel 的三种状态和三种操作结果
操作     空值(nil) 非空已关闭 非空未关闭
关闭     panic     panic   成功关闭
发送数据  永久阻塞   panic    阻塞或者发送成功
接受数据  永久阻塞   永不阻塞  阻塞或者成功接受*/

//关于通道和协程的垃圾回收

/*注意，一个通道被其发送数据协程队列和接收数据协程队列中的所有协程引用着。
因此，如果一个通道的这两个队列只要有一个不为空，则此通道肯定不会被垃圾回收。
另一方面，如果一个协程处于一个通道的某个协程队列之中，则此协程也肯定不会被垃圾回收，
即使此通道仅被此协程所引用。事实上，一个协程只有在退出后才能被垃圾回收。
*/

/*通道关闭原则
一个常用的使用Go通道的原则是不要在数据接收方或者在有多个发送者的情况下关闭通道。换句话说，我们只应该让一个通道唯一的发送者关闭此通道*/

func doCheckClose(taskCh chan int) {
	for {
		select {
		case t, beforeClosed := <-taskCh:
			if !beforeClosed {
				fmt.Println("taskCh has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTaskCheckClose() {
	taskCh := make(chan int, 10)
	go doCheckClose(taskCh)
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	close(taskCh)
}

// 使用sync.Once 或互斥锁确保channel 只被关闭一次

type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		C: make(chan int),
	}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func testMyChannel() {
	myChannel := NewMyChannel()
	fmt.Printf("%v\n", myChannel)
	go doCheckClose(myChannel.C)
	for i := 0; i < 10; i++ {
		myChannel.C <- i
	}
	myChannel.SafeClose()
	myChannel.SafeClose()
	myChannel.SafeClose()
}
