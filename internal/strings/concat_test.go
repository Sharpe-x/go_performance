package strings

import "testing"

// 生成了一个长度为 10 的字符串，并拼接 1w 次。
func benchmarkConcat(b *testing.B, concatFunc func(int, string) string) {
	str := randomString(10)
	for i := 0; i < b.N; i++ {
		concatFunc(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B) {
	benchmarkConcat(b, plusConcat)
}

func BenchmarkSprintfConcat(b *testing.B) {
	benchmarkConcat(b, sprintfConcat)
}

func BenchmarkBuildersConcat(b *testing.B) {
	benchmarkConcat(b, builderConcat)
}

func BenchmarkBufferConcat(b *testing.B) {
	benchmarkConcat(b, bufferConcat)
}

func BenchmarkByteConcat(b *testing.B) {
	benchmarkConcat(b, byteConcat)
}

func BenchmarkPreByteConcat(b *testing.B) {
	benchmarkConcat(b, preByteConcat)
}

func BenchmarkPreBuilderConcat(b *testing.B) {
	benchmarkConcat(b, preBuilderConcat)
}
