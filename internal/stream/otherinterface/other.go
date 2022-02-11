package otherinterface

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// SectionReader 类型 读取数据流中部分数据。

/*type SectionReader struct {
	r     ReaderAt    // 该类型最终的 Read/ReadAt 最终都是通过 r 的 ReadAt 实现
	base  int64        // NewSectionReader 会将 base 设置为 off
	off   int64        // 从 r 中的 off 偏移处开始读取数据
	limit int64        // limit - off = SectionReader 流的长度
}*/

// 常见的创建函数
// func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
// NewSectionReader 返回一个 SectionReader，它从 r 中的偏移量 off 处读取 n 个字节后以 EOF 停止。也就是说，SectionReader 只是内部（内嵌）ReaderAt 表示的数据流的一部分：从 off 开始后的 n 个字节。这个类型的作用是：方便重复操作某一段 (section) 数据流；或者同时需要 ReadAt 和 Seek 的功能。

// LimitedReader 类型
// type LimitedReader struct {
//    R Reader // underlying reader，最终的读取操作通过 R.Read 完成
//    N int64  // max bytes remaining
//}
// 从 R 读取但将返回的数据量限制为 N 字节。每调用一次 Read 都将更新 N 来反应新的剩余数量。也就是说，最多只能返回 N 字节数据。LimitedReader 只实现了 Read 方法（Reader 接口）。

func limitedReaderExample() {
	content := "This is a Example"
	reader := strings.NewReader(content)
	limitedReader := &io.LimitedReader{R: reader, N: 9}
	for limitedReader.N > 0 {
		tmp := make([]byte, 2)
		limitedReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}
}

// PipeReader 和 PipeWriter 类型
// PipeReader（一个没有任何导出字段的 struct）是管道的读取端。它实现了 io.Reader 和 io.Closer 接口。结构定义如下：
/*type PipeReader struct {
	p *pipe
}*/

// 关于 PipeReader.Read 方法的说明：从管道中读取数据。该方法会堵塞，直到管道写入端开始写入数据或写入端被关闭。如果写入端关闭时带有 error（即调用 CloseWithError 关闭），该Read返回的 err 就是写入端传递的error；否则 err 为 EOF。
//
/*PipeWriter（一个没有任何导出字段的 struct）是管道的写入端。它实现了 io.Writer 和 io.Closer 接口。结构定义如下：

type PipeWriter struct {
	p *pipe
}*/
// 关于 PipeWriter.Write 方法的说明：写数据到管道中。该方法会堵塞，直到管道读取端读完所有数据或读取端被关闭。如果读取端关闭时带有 error（即调用 CloseWithError 关闭），该Write返回的 err 就是读取端传递的error；否则 err 为 ErrClosedPipe。
//io.Pipe() 用于创建一个同步的内存管道 (synchronous in-memory pipe)，函数签名：
//func Pipe() (*PipeReader, *PipeWriter)
//它将 io.Reader 连接到 io.Writer。一端的读取匹配另一端的写入，直接在这两端之间复制数据；它没有内部缓存。它对于并行调用 Read 和 Write 以及其它函数或 Close 来说都是安全的。一旦等待的 I/O
//结束，Close 就会完成。并行调用 Read 或并行调用 Write 也同样安全：同种类的调用将按顺序进行控制。
//正因为是同步的，因此不能在一个 goroutine 中进行读和写。
// 读关闭管道
//func (r *PipeReader) Close() error
// 读关闭管道并传入错误信息。
// func (r *PipeReader) CloseWithError(err error) error
// 从管道中读取数据，如果管道被关闭，则会返会一个错误信息：
// 1、如果写入端通过 CloseWithError 方法关闭了管道，则返回关闭时传入的错误信息。
// 2、如果写入端通过 Close 方法关闭了管道，则返回 io.EOF。
// 3、如果是读取端关闭了管道，则返回 io.ErrClosedPipe。

func pipeExample() {
	r, w := io.Pipe()
	// 启用一个 goruntine 进行读取
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			r.CloseWithError(errors.New("管道被读取端关闭"))
			fmt.Printf("read: %d, %v, %s\n", n, err, buf[:n])
		}
	}()

	// 主程序进行写入
	n, err := w.Write([]byte("Hello World !"))
	fmt.Printf("Write: %d,%v\n", n, err)

	time.Sleep(time.Second * 1)

}

func pipeWriterClose() {
	r, w := io.Pipe()
	// 启用一个例程进行读取
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])
		}
	}()
	// 主例程进行写入
	n, err := w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)

	w.CloseWithError(errors.New("管道被写入端关闭"))
	n, err = w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)
	time.Sleep(time.Second * 1)
}

func pipeExample2() {
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(30 * time.Second)
}

func PipeWrite(writer *io.PipeWriter) {
	data := []byte("Go语言中文网")
	for i := 0; i < 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节 %d\n", n)
	}
	writer.CloseWithError(errors.New("写入段已关闭"))
}

func PipeRead(reader *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接口端开始阻塞5秒钟...")
		time.Sleep(5 * time.Second)
		fmt.Println("接收端开始接受")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字节: %d\n buf内容: %s\n", n, buf)
	}
}
