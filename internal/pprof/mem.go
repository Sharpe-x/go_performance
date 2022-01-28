package pprof

import (
	"github.com/pkg/profile"
	"math/rand"
	"strings"
)

// 内存性能分析(Memory profiling) 记录堆内存分配时的堆栈信息，忽略栈内存分配信息。
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func concat2(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()
}

func GetMemProfile() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(100)
}

func GetMemProfile2() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat2(100)
}
