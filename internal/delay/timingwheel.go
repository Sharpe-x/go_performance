package main

import (
	"container/list"
	"fmt"
	"github.com/pkg/errors"
	"go-performance/internal/delay/delayqueue"
	"go-performance/internal/delay/utils"
	"log"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 简单时间轮 https://www.luozhiyun.com/archives/444
// 在时间轮中存储任务的是一个环形队列，底层采用数组实现 数组中的每个元素可以存放一个定时任务列表 定时任务列表是一个环形的双向链表 链表中的每一项都是定时任务项， 封装了真正的定时任务
// 时间轮有多个时间格组成 每个时间格代表当前时间轮的基本时间跨度（tickMS）. 时间轮的时间格个数是固定的，可用wheelSize来表示，那么整个时间轮的基本时间跨度（interval） 可以通过公式
// tickMs * wheelSize 计算得出
// 时间轮还有一个表盘指针（currentTime），用来表示时间轮当前所处的时间，currentTime 是 tickMs 的整数倍。currentTime指向的地方是表示到期的时间格，表示需要处理的时间格所对应的链表中的所有任务。
// 初始情况下表盘指针 currentTime 指向时间格0，若时间轮的 tickMs 为 1ms 且 wheelSize 等于10，那么interval则等于10s。
//如下图此时有一个定时为2s的任务插进来会存放到时间格为2的任务链表中，用红色标记。随着时间的不断推移，指针 currentTime 不断向前推进，如果过了2s，那么 currentTime 会指向时间格2的位置，会将此时间格的任务链表获取出来处理。
// 如果当前的指针 currentTime 指向的是2，此时如果插入一个9s的任务进来，那么新来的任务会服用原来的时间格链表，会存放到时间格1中
// 这里所讲的时间轮都是简单时间轮，只有一层，总体时间范围在 currentTime 和 currentTime+interval 之间。如果现在有一个15s的定时任务是需要重新开启一个时间轮，
//设置一个时间跨度至少为15s的时间轮才够用。但是这样扩充是没有底线的，如果需要一个1万秒的时间轮，那么就需要一个这么大的数组去存放，不仅占用很大的内存空间，而且也会因为需要遍历这么大的数组从而拉低效率。
//因此引入了层级时间轮的概念。

// 层级时间轮
// 如图是一个两层的时间轮，第二层时间轮也是由10个时间格组成，每个时间格的跨度是10s。第二层的时间轮的 tickMs 为第一层时间轮的 interval，即10s。每一层时间轮的 wheelSize 是固定的，都是10，那么第二层的时间轮的总体时间跨度 interval 为100s。
// 图中展示了每个时间格对应的过期时间范围， 我们可以清晰地看到， 第二层时间轮的第0个时间格的过期时间范围是 [0,9]。也就是说, 第二层时间轮的一个时间格就可以表示第一层时间轮的所有(10个)时间格；
//如果向该时间轮中添加一个15s的任务，那么当第一层时间轮容纳不下时，进入第二层时间轮，并插入到过期时间为[10，19]的时间格中。
// 随着时间的流逝，当原本15s的任务还剩下5s的时候，这里就有一个时间轮降级的操作，此时第一层时间轮的总体时间跨度已足够，此任务被添加到第一层时间轮到期时间为5的时间格中，之后再经历5s后，此任务真正到期，最终执行相应的到期操作。

type TimingWheel struct {
	// 时间跨度 单位是毫秒
	tick int64
	// 时间轮个数
	wheelSize int64
	// 总跨度
	interval int64
	// 当前指针指向时间
	currentTime int64
	// 时间格列表
	buckets []*bucket
	// 延迟队列
	queue *delayqueue.DelayQueue
	// 上级时间轮引用
	overflowWheel unsafe.Pointer

	exitC     chan struct{}
	waitGroup utils.WaitGroupWrapper
}

func NewTimingWheel(tick time.Duration, wheelSize int64) *TimingWheel {
	// 将传入的tick转化成毫秒
	tickMs := int64(tick / time.Millisecond)
	// 如果小于零，那么panic
	if tickMs <= 0 {
		panic(errors.New("tick must be greater than or equal to 1ms"))
	}
	// 设置开始时间
	startMs := utils.TimeToMs(time.Now().UTC())
	// 初始化TimingWheel
	return newTimingWheel(tickMs, wheelSize, startMs, delayqueue.New(wheelSize))
}

func newTimingWheel(tickMs int64, wheelSize int64, startMs int64, queue *delayqueue.DelayQueue) *TimingWheel {
	// 初始化buckets 的大小
	buckets := make([]*bucket, wheelSize)
	for i := range buckets {
		buckets[i] = newBucket()
	}

	return &TimingWheel{
		tick:        tickMs,
		wheelSize:   wheelSize,
		currentTime: utils.Truncate(startMs, tickMs),
		interval:    tickMs * wheelSize,
		buckets:     buckets,
		queue:       queue,
		exitC:       make(chan struct{}),
	}
}

//Start方法会启动两个goroutines。第一个goroutines用来调用延迟队列的queue的Poll方法，这个方法会一直循环获取队列里面的数据，然后将到期的数据放入到queue的C管道中；
//第二个goroutines 会无限循环获取queue中C的数据，如果C中有数据表示已经到期，那么会先调用advanceClock方法将当前时间 currentTime 往前移动到 bucket的到期时间，然后再调用Flush方法取出bucket中的队列，并调用addOrRun方法执行。

func (tw *TimingWheel) Start() {
	// Poll会执行一个无限循环，将到期的元素放入到queue的C管道中
	tw.waitGroup.Wrap(func() {
		tw.queue.Poll(tw.exitC, func() int64 {
			return utils.TimeToMs(time.Now().UTC())
		})
	})

	// 开启无限循环获取queue中C的数据
	tw.waitGroup.Wrap(func() {
		for {
			select {
			// 从队列里面出来的数据都是到期的bucket
			case elem := <-tw.queue.C:
				log.Println("elem is ", elem)
				b := elem.(*bucket)
				// 时间轮会将当前时间 currentTime 往前移动到 bucket的到期时间
				tw.advanceClock(b.Expiration())
				// 取出bucket队列的数据，并调用addOrRun方法执行
				b.Flush(tw.addOrRun)
			case <-tw.exitC:
				return
			}
		}
	})
}

// advanceClock方法会根据到期时间来从新设置currentTime，从而推进时间轮前进。
func (tw *TimingWheel) advanceClock(expiration int64) {
	currentTime := atomic.LoadInt64(&tw.currentTime)
	// 过期时间大于等于（当前时间+tick）
	if expiration >= currentTime+tw.tick {
		// 将currentTime设置为expiration，从而推进currentTime
		currentTime = utils.Truncate(expiration, tw.tick)
		atomic.StoreInt64(&tw.currentTime, currentTime)
		// 如果有上层时间轮，那么递归调用上层时间轮的引用
		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel != nil {
			(*TimingWheel)(overflowWheel).advanceClock(currentTime)
		}
	}
}

func (tw *TimingWheel) addOrRun(t *Timer) {
	// 如果已经过期，那么直接执行
	if !tw.add(t) {
		go t.task()
	}
}

// add方法根据到期时间来分成了三部分，第一部分是小于当前时间+tick，表示已经到期，那么返回false执行任务即可；
// 第二部分的判断会根据expiration是否小于时间轮的跨度，如果小于的话表示该定时任务可以放入到当前时间轮中，通过取模找到buckets对应的时间格并放入到bucket队列中，
//	SetExpiration方法会根据传入的参数来判断是否已经执行过延迟队列的Offer方法，防止重复插入；
// 第三部分表示该定时任务的时间跨度超过了当前时间轮，需要升级到上一层的时间轮中。需要注意的是，上一层的时间轮的tick是当前时间轮的interval，延迟队列还是同一个，然后设置为指针overflowWheel，并调用add方法往上层递归。
// 使用了DelayQueue加环形队列的方式实现了时间轮。对定时任务项的插入和删除操作而言，TimingWheel时间复杂度为 O(1)，
//在DelayQueue中的队列使用的是优先队列，时间复杂度是O(log n)，但是由于buckets列表实际上是非常小的，所以并不会影响性能。

func (tw *TimingWheel) add(t *Timer) bool {
	currentTime := atomic.LoadInt64(&tw.currentTime)
	// 已经过期
	if t.expiration < currentTime+tw.tick {
		return false
	} else if t.expiration < currentTime+tw.interval { // 到期时间在第一层环内
		// 获取时间轮的位置
		virtualID := t.expiration / tw.tick
		b := tw.buckets[virtualID%tw.wheelSize]
		// 将任务放入到bucket队列中
		b.Add(t)
		// 如果是相同的时间，那么返回false，防止被多次插入到队列中
		if b.SetExpiration(virtualID * tw.tick) {
			log.Println("insert b ", b, b.expiration)
			tw.queue.Offer(b, b.expiration)
		}
		return true
	} else {
		// 如果放入的到期时间超过第一层时间轮，那么放到上一层中去
		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel == nil {
			atomic.CompareAndSwapPointer(
				// 这里tick变成了interval
				&tw.overflowWheel, nil, unsafe.Pointer(newTimingWheel(tw.interval, tw.wheelSize, currentTime, tw.queue)))
			overflowWheel = atomic.LoadPointer(&tw.overflowWheel)
		}
		// 往上递归
		return (*TimingWheel)(overflowWheel).add(t)
	}
}

// Timer Timer是时间轮的最小执行单元，是定时任务的封装，到期后会调用task来执行任务。
type Timer struct {
	// 到期时间
	expiration int64
	// 要被执行的具体任务
	task func()
	// Timer 所在bucket 的指针
	b unsafe.Pointer
	// bucket 列表中对应的元素
	element *list.Element
}

func (t *Timer) getBucket() *bucket {
	return (*bucket)(atomic.LoadPointer(&t.b))
}

func (t *Timer) setBucket(b *bucket) {
	atomic.StorePointer(&t.b, unsafe.Pointer(b))
}

// Stop prevents the Timer from firing. It returns true if the call
// stops the timer, false if the timer has already expired or been stopped.
//
// If the timer t has already expired and the t.task has been started in its own
// goroutine; Stop does not wait for t.task to complete before returning. If the caller
// needs to know whether t.task is completed, it must coordinate with t.task explicitly.
func (t *Timer) Stop() bool {
	stopped := false
	for b := t.getBucket(); b != nil; b = t.getBucket() {
		// If b.Remove is called just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// this may fail to remove t due to the change of t's bucket.
		stopped = b.Remove(t)

		// Thus, here we re-get t's possibly new bucket (nil for case 1, or ab (non-nil) for case 2),
		// and retry until the bucket becomes nil, which indicates that t has finally been removed.
	}
	return stopped
}

type bucket struct {
	// 任务的过期时间
	expiration int64

	mu sync.Mutex
	// 相同过期时间的任务队列
	timers *list.List
}

func newBucket() *bucket {
	return &bucket{
		timers:     list.New(),
		expiration: -1,
	}
}

func (b *bucket) Flush(reinsert func(*Timer)) {
	var ts []*Timer
	b.mu.Lock()
	// 循环获取bucket队列节点
	for e := b.timers.Front(); e != nil; {
		next := e.Next()
		t := e.Value.(*Timer)
		// 将头节点移除
		b.remove(t)
		ts = append(ts, t)
		e = next
	}
	b.mu.Unlock()
	b.SetExpiration(-1)
	for _, t := range ts {
		reinsert(t)
	}
}

func (b *bucket) Expiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}

func (b *bucket) SetExpiration(expiration int64) bool {
	return atomic.SwapInt64(&b.expiration, expiration) != expiration
}

func (b *bucket) Add(t *Timer) {
	b.mu.Lock()
	defer b.mu.Unlock()
	e := b.timers.PushBack(t)
	t.setBucket(b)
	t.element = e
}

func (b *bucket) remove(t *Timer) bool {
	if t.getBucket() != b {
		// If remove is called from t.Stop, and this happens just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// then t.getBucket will return nil for case 1, or ab (non-nil) for case 2.
		// In either case, the returned value does not equal to b.
		return false
	}
	b.timers.Remove(t.element)
	t.setBucket(nil)
	t.element = nil
	return true
}

func (b *bucket) Remove(t *Timer) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.remove(t)
}

func (tw *TimingWheel) AfterFunc(d time.Duration, f func()) *Timer {
	t := &Timer{
		expiration: utils.TimeToMs(time.Now().UTC().Add(d)),
		task:       f,
	}
	tw.addOrRun(t)
	return t
}

func (tw *TimingWheel) Stop() {
	close(tw.exitC)
	tw.waitGroup.Wait()
}

func main() {
	tw := NewTimingWheel(time.Second, 5)
	tw.Start()
	defer tw.Stop()
	// 添加任务
	tw.AfterFunc(time.Second*60, func() {
		fmt.Println("The Timer fires")
	})

	time.Sleep(time.Second * 100)
}
