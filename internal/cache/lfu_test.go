package main

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lfu := Constructor(3)
	fmt.Println(lfu.Get(1))
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)
	lfu.Put(1, 11)
	//fmt.Println(lfu.Get(1))
	//fmt.Println(lfu.Get(2))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(2))
}
