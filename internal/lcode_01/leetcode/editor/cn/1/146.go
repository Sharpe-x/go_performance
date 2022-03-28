//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
//
// 实现 LRUCache 类：
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
// Related Topics 设计 哈希表 链表 双向链表 👍 1987 👎 0

package main

import (
	"container/list"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	capacity int
	list     *list.List
	KeyMaps  map[int]*list.Element
}

type entry struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		capacity: capacity,
		KeyMaps:  make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.KeyMaps[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.KeyMaps[key]; ok {
		// update value
		elem.Value.(*entry).val = value
		// 移到队头
		this.list.MoveToFront(elem)
	} else {
		et := this.list.PushFront(&entry{
			key: key,
			val: value,
		})
		this.KeyMaps[key] = et
	}

	// 如果大于capacity 删除尾部
	if len(this.KeyMaps) > this.capacity {
		et := this.list.Back()
		fmt.Println("delete ", et.Value.(*entry).key)
		delete(this.KeyMaps, et.Value.(*entry).key)
		this.list.Remove(et)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	lru := Constructor(2)
	//fmt.Println(lru.Get(1))
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	//	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	/*	fmt.Println(lru.Get(1))
		fmt.Println(lru.Get(3))
		fmt.Println(lru.Get(4))*/
}
