package main

import (
	"container/list"
	"hash/fnv"
	"sync"
)

type command int

const (
	MoveToFront command = iota
	PushFront
	Delete
)

type clear struct {
	done chan struct{}
}

type CLRUCache struct {
	sync.RWMutex
	cap         int
	list        *list.List
	buckets     []*bucket
	bucketMask  uint32
	deletePairs chan *list.Element
	movePairs   chan *list.Element
	control     chan interface{}
}

type Pair struct {
	key   string
	value interface{}
	cmd   command
}

func New(capacity int) *CLRUCache {
	c := &CLRUCache{
		cap:        capacity,
		list:       list.New(),
		bucketMask: uint32(1024) - 1,
		buckets:    make([]*bucket, 1024),
	}
	for i := 0; i < 1024; i++ {
		c.buckets[i] = &bucket{
			keys: make(map[string]*list.Element),
		}
	}
	c.restart()
	return c
}

// Get define
func (c *CLRUCache) Get(key string) interface{} {
	el := c.bucket(key).get(key)
	if el == nil {
		return nil
	}
	c.move(el)
	return el.Value.(Pair).value
}

// Put define
func (c *CLRUCache) Put(key string, value interface{}) {
	el, exist := c.bucket(key).set(key, value)
	if exist != nil {
		c.deletePairs <- exist
	}
	c.move(el)
}

func (c *CLRUCache) move(el *list.Element) {
	select {
	case c.movePairs <- el:
	default:
	}
}

// Delete define
func (c *CLRUCache) Delete(key string) bool {
	el := c.bucket(key).delete(key)
	if el != nil {
		c.deletePairs <- el
		return true
	}
	return false
}

// Clear define
func (c *CLRUCache) Clear() {
	done := make(chan struct{})
	c.control <- clear{done: done}
	<-done
}

// Count define
func (c *CLRUCache) Count() int {
	count := 0
	for _, b := range c.buckets {
		count += b.pairCount()
	}
	return count
}

func (c *CLRUCache) stop() {
	close(c.movePairs)
	<-c.control
}

func (c *CLRUCache) restart() {
	c.deletePairs = make(chan *list.Element, 128)
	c.movePairs = make(chan *list.Element, 128)
	c.control = make(chan interface{})
}

func (c *CLRUCache) worker() {
	defer close(c.control)
	for {
		select {
		case el, ok := <-c.movePairs:
			if ok == false {
				goto clean
			}
			if c.doMove(el) && c.list.Len() > c.cap {
				el := c.list.Back()
				c.list.Remove(el)
				c.bucket(el.Value.(Pair).key).delete(el.Value.(Pair).key)
			}
		case el := <-c.deletePairs:
			c.list.Remove(el)
		case control := <-c.control:
			switch msg := control.(type) {
			case clear:
				for _, bucket := range c.buckets {
					bucket.clear()
				}
				c.list = list.New()
				msg.done <- struct{}{}
			}
		}
	}
clean:
	for {
		select {
		case el := <-c.deletePairs:
			c.list.Remove(el)
		default:
			close(c.deletePairs)
			return
		}
	}
}

func (c *CLRUCache) bucket(key string) *bucket {
	h := fnv.New32a()
	h.Write([]byte(key))
	return c.buckets[h.Sum32()&c.bucketMask]
}

func (c *CLRUCache) doMove(el *list.Element) bool {
	if el.Value.(Pair).cmd == MoveToFront {
		c.list.MoveToFront(el)
		return false
	}
	newel := c.list.PushFront(el.Value.(Pair))
	c.bucket(el.Value.(Pair).key).update(el.Value.(Pair).key, newel)
	return true
}

type bucket struct {
	sync.RWMutex
	keys map[string]*list.Element
}

func (b *bucket) pairCount() int {
	b.RLock()
	defer b.RUnlock()
	return len(b.keys)
}

func (b *bucket) get(key string) *list.Element {
	b.RLock()
	defer b.RUnlock()
	if el, ok := b.keys[key]; ok {
		return el
	}
	return nil
}

func (b *bucket) set(key string, value interface{}) (*list.Element, *list.Element) {
	el := &list.Element{Value: Pair{key: key, value: value, cmd: PushFront}}
	b.Lock()
	exist := b.keys[key]
	b.keys[key] = el
	b.Unlock()
	return el, exist
}

func (b *bucket) update(key string, el *list.Element) {
	b.Lock()
	b.keys[key] = el
	b.Unlock()
}

func (b *bucket) delete(key string) *list.Element {
	b.Lock()
	el := b.keys[key]
	delete(b.keys, key)
	b.Unlock()
	return el
}

func (b *bucket) clear() {
	b.Lock()
	defer b.Unlock()
	b.keys = make(map[string]*list.Element)
}
