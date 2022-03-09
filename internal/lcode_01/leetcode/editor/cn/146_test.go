package main

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(2)
	fmt.Println(lru.Get(1))
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	fmt.Println(lru.KeyMaps)
}
