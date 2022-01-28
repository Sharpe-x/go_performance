package benchmark

import (
	"testing"
	"time"
)

// BenchmarkFib 基准测试
// go test 命令默认不运行 benchmark 用例的，如果想运行 benchmark 用例，则需要加上 -bench 参数。
// go test -bench .
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ { //b.N 表示这个用例需要运行的次数。b.N 对于每个用例都是不一样的。 b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快。
		Fib(30)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(1000000)
	}
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateWithCap(1000000)
	}
}

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B) {
	benchmarkGenerate(1000, b)
}
func BenchmarkGenerate10000(b *testing.B) {
	benchmarkGenerate(10000, b)
}
func BenchmarkGenerate100000(b *testing.B) {
	benchmarkGenerate(100000, b)
}
func BenchmarkGenerate1000000(b *testing.B) {
	benchmarkGenerate(1000000, b)
}

func BenchmarkFibResetTimer(b *testing.B) {
	time.Sleep(time.Second * 3) //耗时准备任务 比如说数据准备

	b.ResetTimer() // 重置计时器

	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}

/*func BenchmarkFibUseStartAndEndTimer(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		time.Sleep(time.Second * 1) //耗时准备任务 比如说数据清理和重置
		b.StartTimer()
		Fib(30)
	}
}*/
