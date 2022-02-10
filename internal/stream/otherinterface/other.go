package otherinterface

import (
	"fmt"
	"io"
	"strings"
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
