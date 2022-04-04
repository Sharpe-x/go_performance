package main

import "container/heap"

// LFU 的另外一个思路是利用 Index Priority Queue 这个数据结构。别被名字吓到，Index Priority Queue = map + Priority Queue，仅此而已。
//利用 Priority Queue 维护一个最小堆，堆顶是访问次数最小的元素。map 中的 value 存储的是优先队列中结点。

type LFUCacheByIPQ struct {
	capacity int
	pq       PriorityQueue
	hash     map[int]*Item
	counter  int
}

func ConstructorLFUCacheByIPQ(capacity int) LFUCacheByIPQ {
	return LFUCacheByIPQ{
		capacity: capacity,
		pq:       PriorityQueue{},
		hash:     make(map[int]*Item, capacity),
	}
}

type PriorityQueue []*Item
type Item struct {
	// count 值用来决定哪个是最老的元素
	// index 值用来 re-heapify 调整堆的。
	key, value, frequency, count, index int
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, frequency int, count int) {
	item.value = value
	item.count = count
	item.frequency = frequency
	heap.Fix(pq, item.index)
}

func (lfu *LFUCacheByIPQ) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}
	if item, ok := lfu.hash[key]; ok {
		lfu.counter++
		lfu.pq.update(item, item.value, item.frequency+1, lfu.counter)
		return item.value
	}
	return -1
}

func (lfu *LFUCacheByIPQ) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	lfu.counter++

	if item, ok := lfu.hash[key]; ok {
		lfu.pq.update(item, value, item.frequency+1, lfu.counter)
		return
	}

	if len(lfu.pq) == lfu.capacity {
		item := heap.Pop(&lfu.pq).(*Item)
		delete(lfu.hash, item.key)
	}

	item := &Item{
		key:   key,
		value: value,
		count: lfu.counter,
	}
	heap.Push(&lfu.pq, item)
	lfu.hash[key] = item
}
