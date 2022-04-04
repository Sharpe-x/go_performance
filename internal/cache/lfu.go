package main

import "container/list"

// LFU 是 Least Frequently Used 的缩写，即最不经常最少使用，也是一种常用的页面置换算法，选择访问计数器最小的页面予以淘汰。
// 比 LRU 特别的地方。如果淘汰的页面访问次数有多个相同的访问次数，选择最靠尾部的。(更旧的)
// 还有 1 个问题需要考虑，一个是如何按频次排序？相同频次，按照先后顺序排序。如果你开始考虑排序算法的话，思考方向就偏离最佳答案了
//。排序至少 O(nlogn)。重新回看 LFU 的工作原理，会发现它只关心最小频次。其他频次之间的顺序并不关心。所以不需要排序。
//用一个 min 变量保存最小频次，淘汰时读取这个最小值能找到要删除的结点。相同频次按照先后顺序排列，这个需求还是用双向链表实现，双向链表插入的顺序体现了结点的先后顺序。
//相同频次对应一个双向链表，可能有多个相同频次，所以可能有多个双向链表。
//用一个 map 维护访问频次和双向链表的对应关系。删除最小频次时，通过 min 找到最小频次，
//然后再这个 map 中找到这个频次对应的双向链表，在双向链表中找到最旧的那个结点删除。这就解决了 LFU 删除操作。

type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type entry struct {
	key, value, frequently int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	elem, ok := lfu.nodes[key] // key 不存在直接返回
	if !ok {
		return -1
	}

	curEntry := elem.Value.(*entry)
	lfu.lists[curEntry.frequently].Remove(elem) //因为已经使用了，所以从frequently 维护的链表移除，frequently++
	curEntry.frequently++

	if _, ok = lfu.lists[curEntry.frequently]; !ok { // 没有frequently次使用的队列 建一个
		lfu.lists[curEntry.frequently] = list.New()
	}
	freList := lfu.lists[curEntry.frequently]

	node := freList.PushFront(curEntry) // freList 将最近使用entry的放到头部
	lfu.nodes[key] = node               //更新key 对应的节点

	// 老的 frequency 对应的双向链表中是否已经为空，如果空了，min++。
	if curEntry.frequently-1 == lfu.min && lfu.lists[curEntry.frequently-1].Len() == 0 {
		lfu.min++
	}
	return curEntry.value
}

func (lfu *LFUCache) Put(key, value int) {
	if lfu.capacity == 0 {
		return
	}

	// 如果存在 更新访问次数
	if elem, ok := lfu.nodes[key]; ok {
		curEntry := elem.Value.(*entry)
		curEntry.value = value
		// 更新逻辑和 Get 操作一致。
		lfu.Get(key)
		return
	}

	// 不存在且缓存已满 需要删除
	if lfu.capacity == len(lfu.nodes) {
		curList := lfu.lists[lfu.min]                  // 拿到最少访问次数的链表
		backElem := curList.Back()                     // 拿到最旧的一个
		delete(lfu.nodes, backElem.Value.(*entry).key) //删除
		curList.Remove(backElem)                       //链表中也删除
	}

	// 新插入的页面访问次数一定为 1，所以 min 此时置为 1
	lfu.min = 1
	curEntry := &entry{
		key:        key,
		value:      value,
		frequently: 1,
	}
	if _, ok := lfu.lists[1]; !ok {
		lfu.lists[1] = list.New()
	}
	lfu.nodes[key] = lfu.lists[1].PushFront(curEntry)
}
