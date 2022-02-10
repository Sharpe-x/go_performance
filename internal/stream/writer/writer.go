package writer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*type Writer interface {
	//Write 方法有两个返回值，一个是写入到目标资源的字节数，一个是发生错误时的错误。
	Write(p []byte) (n int, err error)
}*/

// Writer 表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源
// 实现这个接口就需要实现如下的功能
// Write 将len（p） 个字节写入到基本数据流中。它返回从p中被写入的字节数 （0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
// 若Write 返回的n < len(P) ,他就必须返回一个非nil 的错误

func writeExample() {

	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	var writer bytes.Buffer
	for _, proverb := range proverbs {
		n, err := writer.Write([]byte(proverb))
		if err != nil {
			fmt.Println("failed to write data", err)
			os.Exit(1)
		}
		if n != len(proverb) {
			fmt.Println("failed to write ", n, len(proverbs))
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
}

// chanWriter 实现 io.Writer ，它将其内容作为字节序列写入 channel 。
type chanWriter struct {
	// ch 实际上就是目标资源
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{
		make(chan byte, 1024),
	}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

// Closer 接口包装了基本的 Close 方法，用于关闭数据读写。Close 一般用于关闭文件，关闭通道，关闭连接，关闭数据库等，在不同的标准库实现中实现。
// type Closer interface {
//    Close() error
//}

// Close chanWriter 实现了接口 io.Closer ，调用方法 writer.Close() 来正确地关闭channel，以避免发生泄漏和死锁。
func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func testChanWriter() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me"))
	}()

	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// seeker
// Seeker 接口包装了基本的 Seek 方法，用于移动数据的读写指针。
// type Seeker interface {
//    Seek(offset int64, whence int) (ret int64, err error)
//}
// Seek 设置下一次读写操作的指针位置，每次的读写操作都是从指针位置开始的。
// whence 的含义：
// 如果 whence 为 0：表示从数据的开头开始移动指针。
// 如果 whence 为 1：表示从数据的当前指针位置开始移动指针。
// 如果 whence 为 2：表示从数据的尾部开始移动指针。
// offset 是指针移动的偏移量。 返回新指针位置和遇到的错误。
// whence 的值，在 io 包中定义了相应的常量，应该使用这些常量
// const (
//  SeekStart   = 0 // seek relative to the origin of the file
//  SeekCurrent = 1 // seek relative to the current offset
//  SeekEnd     = 2 // seek relative to the end
//)

// 组合接口 这些接口的作用是：有些时候同时需要某两个接口的所有功能，
// 即必须同时实现了某两个接口的类型才能够被传入使用。可见，io 包中有大量的“小接口”，这样方便组合为“大接口”。
/*type ReadSeeker interface {
	Reader
	Seeker
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}*/

// 其他接口
// ReaderFrom 接口包装了基本的 ReadFrom 方法，用于从 r 中读取数据存入自身。直到遇到 EOF 或读取出错为止，返回读取的字节数和遇到的错误。
/*type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}*/
// 需要实现接口的功能
// ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。
// 如果 ReaderFrom 可用，Copy 函数就会使用它。

func readFromExample() {
	file, err := os.Open("writer_test.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	_, err = writer.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}

// WriterTo 接口包装了基本的 WriteTo 方法，用于将自身的数据写入 w 中。直到数据全部写入完毕或遇到错误为止，返回写入的字节数和遇到的错误。
// type WriterTo interface {
//    WriteTo(w Writer) (n int64, err error)
//}
// 需要实现接口的功能
//  WriteTo 将数据写入 w 中，直到没有数据可写或发生错误。其返回值 n 为写入的字节数。 在写入过程中遇到的任何错误也将被返回。
// 如果 WriterTo 可用，Copy 函数就会使用它。
// ReaderFrom 和 WriterTo 接口的方法接收的参数是 io.Reader 和 io.Writer 类型。

// ReaderAt 接口包装了基本的 ReadAt 方法，用于将自身的数据写入 p 中。ReadAt 忽略之前的读写位置，从起始位置的 off 偏移处开始读取。
// type ReaderAt interface {
//    ReadAt(p []byte, off int64) (n int, err error)
//}
// ReadAt 从基本输入源的偏移量 off 处开始，将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)）以及任何遇到的错误。
//当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节。在这一点上，ReadAt 比 Read 更严格。
//即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。若可读取的数据不到 len(p) 字节，ReadAt 就会阻塞,直到所有数据都可用或一个错误发生。 在这一点上 ReadAt 不同于 Read。
//若 n = len(p) 个字节从输入源的结尾处由 ReadAt 返回，Read可能返回 err == EOF 或者 err == nil
//若 ReadAt 携带一个偏移量从输入源读取，ReadAt 应当既不影响偏移量也不被它所影响。
//可对相同的输入源并行执行 ReadAt 调用。

func readAtExample() {
	reader := strings.NewReader("Hello Golang and Rust")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d", p, n)
}

// WriterAt 接口包装了基本的 WriteAt 方法，用于将 p 中的数据写入自身。ReadAt 忽略之前的读写位置，从起始位置的 off 偏移处开始写入。
// type WriterAt interface {
//    WriteAt(p []byte, off int64) (n int, err error)
//}
// 需要实现接口的功能
//
//WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off 处的基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。若 WriteAt 返回的 n < len(p)，它就必须返回一个 非nil 的错误。
//若 WriteAt 携带一个偏移量写入到目标中，WriteAt 应当既不影响偏移量也不被它所影响。
//若被写区域没有重叠，可对相同的目标并行执行 WriteAt 调用。

func writerAtExample() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("123456790")
	n, err := file.WriteAt([]byte("8910"), 7) // 会覆盖该位置的内容
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
