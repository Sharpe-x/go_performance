package reader

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// stream就是数据流，数据流的概念其实非常基础，最早是在通讯领域使用的概念，这个概念最初在 1998 年由 Henzinger 在文献中提出，
//他将数据流定义为 “只能以事先规定好的顺序被读取一次的数据的一个序列”
//数据流就是由数据形成的流，就像由水形成的水流，非常形象，现代语言中，基本上都会有流的支持，比如 C++ 的 iostream，Node.js 的 stream 模块，以及 golang 的 io 包。
//Stream in Golang与流密切相关的就是 bufio io io/ioutil 这几个包：
// io 为 IO 原语（I/O primitives）提供基本的接口
// io/ioutil 封装一些实用的 I/O 函数
// fmt 实现格式化 I/O，类似 C 语言中的 printf 和 scanf
// bufio 实现带缓冲I/O

// io 包为 I/O 原语提供了基本的接口。在 io 包中最重要的是两个接口：Reader 和 Writer 接口。

//读取器接口

/*type Reader interface {
	// Read 方法有两个返回值，一个是读取到的字节数，一个是发生错误时的错误。如果资源内容已全部读取完毕，应该返回 io.EOF 错误。
	Read(p []byte) (n int, err error)
}*/

/*
	io.Reader 表示一个读取器，它将数据从某个资源读取到传输缓冲区p。在缓冲区中，数据可以被流式传输和使用。
	实现这个接口需要实现如下功能
	Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。
		即使 Read 返回的 n < len(p)，它也会在调用过程中占用 len(p) 个字节作为暂存空间。
		若可读取的数据不到 len(p) 个字节，Read 会返回可用数据，而不是等待更多数据。
		当读取的时候没有数据也没有EOF的时候，会阻塞在这边等待。
	当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file)，它会返回读取的字节数。
		它可能会同时在本次的调用中返回一个non-nil错误,或在下一次的调用中返回这个错误（且 n 为 0）。
		一般情况下, Reader会返回一个非0字节数n, 若 n = len(p) 个字节从输入源的结尾处由 Read 返回，Read可能返回 err == EOF 或者 err == nil。并且之后的 Read() 都应该返回 (n:0, err:EOF)。
    调用者在考虑错误之前应当首先处理返回的数据。这样做可以正确地处理在读取一些字节后产生的 I/O 错误，同时允许EOF的出现。
*/

//对于要用作读取器的类型，它必须实现 io.Reader 接口的唯一一个方法 Read(p []byte)。换句话说，只要实现了 Read(p []byte) ，那它就是一个读取器，使用标准库中已经实现的读写器，来举例

func stringReader() {
	reader := strings.NewReader("Clear is better chan clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(n, string(p[:n]))
	}
}

func alphaReaderCase() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf(string(p[:n]))
	}
	fmt.Println()
}

// 自定义Reader 从流中过滤掉非字母字符。
type alphaReader struct {
	src string // 资源
	cur int    // 当前读到的位置
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{
		src: src,
	}
}

func alpha(r byte) byte {

	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}

	return 0
}

func (a *alphaReader) Read(p []byte) (n int, err error) {
	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	// x 是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		// 每次读取一个字节，执行过滤函数
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}
