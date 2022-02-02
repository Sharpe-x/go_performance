package concurrency

import (
	"encoding/json"
	"sync"
)

// sync.Pool 的使用场景
// 保存和复用临时对象 减少内存分配 降低GC压力

var buf, _ = json.Marshal(Student{Name: "Test", Age: 25})

type Student struct {
	Name   string
	Age    int
	Remark [1024]byte
}

// 声明对象池
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}
