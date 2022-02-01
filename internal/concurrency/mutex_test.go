package concurrency

import (
	"sync"
	"testing"
)

// 单位读写操作时间下降后，读写锁的性能优势下降到
// 因加锁而阻塞的时间占比减小，互斥锁带来的损耗自然就减小了。

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < read*100; j++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}

		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// 读多写少
func BenchmarkReadMore(b *testing.B) {
	benchmark(b, &Lock{}, 9, 1)
}

// // 读多写少
func BenchmarkReadMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 9, 1)
}

// 读少写多
func BenchmarkWriteMore(b *testing.B) {
	benchmark(b, &Lock{}, 1, 9)
}

// 读少写多
func BenchmarkWriteMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 1, 9)
}

// 读写均匀
func BenchmarkEqual(b *testing.B) {
	benchmark(b, &Lock{}, 5, 5)
}

// 读写均匀
func BenchmarkEqualRW(b *testing.B) {
	benchmark(b, &RWLock{}, 5, 5)
}
