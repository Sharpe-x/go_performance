package concurrency

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := studentPool.Get().(*Student)
		_ = json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
