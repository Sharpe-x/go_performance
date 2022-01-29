package strings

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

// 在 Go 语言中，字符串(string) 是不可变的，拼接字符串事实上是创建了一个新的字符串对象。如果代码中存在大量的字符串拼接，对性能会产生严重的影响。

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"

//randomString 生成长度为 n 的随机字符串的函数
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 常见字符串拼接函数

// plusConcat 使用 + 拼接字符串
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

//sprintfConcat 使用 fmt.Sprintf 拼接字符串
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// builderConcat 使用 bytes.Buffer 拼接字符串
func builderConcat(n int, str string) string {
	var builder strings.Builder

	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}

	return builder.String()
}

// bufferConcat 使用 bytes.Buffer 拼接字符串
func bufferConcat(n int, str string) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

// byteConcat 使用 []byte 拼接字符串
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// preByteConcat 使用 []byte 拼接字符串 如果长度是可预知的，那么创建 []byte 时，我们还可以预分配切片的容量(cap)。
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, len(str)*n)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// preBuilderConcat 长度是可预知的，builder预分配切片的容量(cap)。
func preBuilderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
