package delayqueue

import (
	"container/heap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// The start of PriorityQueue implementation.
// Borrowed from https://github.com/nsqio/nsq/blob/master/internal/pqueue/pqueue.go

type item struct {
	Value    interface{}
	Priority int64
	Index    int
}

// this is a priority queue as implemented by a min heap
// the 0th element is the *lowest* value

type priorityQueue []*item

func newPriorityQueue(capacity int64) priorityQueue {
	return make(priorityQueue, 0, capacity)
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(priorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	*pq = (*pq)[0 : n+1]
	it := x.(*item)
	it.Index = n
	(*pq)[n] = it
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	if n < (c/2) && c > 25 {
		npq := make(priorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	it := (*pq)[n-1]
	it.Index = -1
	*pq = (*pq)[0 : n-1]
	return it
}

func (pq *priorityQueue) PeekAndShift(max int64) (*item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}
	it := (*pq)[0]
	if it.Priority > max {
		return nil, it.Priority - max
	}
	heap.Remove(pq, 0)
	return it, 0
}

// The end of PriorityQueue implementation.

// DelayQueue is an unbounded blocking queue of *Delayed* elements, in which
// an element can only be taken when its delay has expired. The head of the
// queue is the *Delayed* element whose delay expired furthest in the past.

type DelayQueue struct {
	C chan interface{}

	mu sync.Mutex
	pq priorityQueue

	sleeping int32 // Similar to the sleeping state of runtime.timers.

	wakeUpC chan struct{}
}

func New(size int64) *DelayQueue {
	return &DelayQueue{
		C:       make(chan interface{}),
		pq:      newPriorityQueue(size),
		wakeUpC: make(chan struct{}),
	}
}

// Offer inserts the element into the current queue.
func (dq *DelayQueue) Offer(elem interface{}, expiration int64) {
	it := &item{
		Value:    elem,
		Priority: expiration,
	}

	dq.mu.Lock()
	heap.Push(&dq.pq, it)
	index := it.Index
	dq.mu.Unlock()

	if index == 0 {
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeUpC <- struct{}{}
		}
	}
}

// Poll starts an infinite loop, in which it continually waits for an element
//
func (dq *DelayQueue) Poll(exitC chan struct{}, nowF func() int64) {
	for {
		now := nowF()
		dq.mu.Lock()
		it, delta := dq.pq.PeekAndShift(now)
		if it == nil {
			// No items left or at least one item is pending.

			// We must ensure the atomicity of the whole operation, which is
			// composed of the above PeekAndShift and the following StoreInt32,
			// to avoid possible race conditions between Offer and Poll.
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()
		if it == nil {
			log.Println("Poll get it is nil and delta is ", delta)
			if delta == 0 {
				// No items left.
				select {
				case <-dq.wakeUpC:
					// wait until a new item is added
					continue
				case <-exitC:
					goto exit
				}
			} else if delta > 0 {
				// At least one item is pending
				select {
				case <-dq.wakeUpC:
					// A new item with an "earlier" expiration than the current "earliest" one is added.
					continue
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeUpC
					}
					continue
				case <-exitC:
					goto exit
				}
			}
		}

		select {
		case dq.C <- it.Value:
			log.Println("写入dq.C ", it.Value, " ", delta)
			// The expired element has been sent out successfully.
		case <-exitC:
			goto exit
		}
	}
exit:
	atomic.StoreInt32(&dq.sleeping, 0)
}
