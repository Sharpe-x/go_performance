package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteString
// WriteString 将字符串 s 写入到 w 中，返回写入的字节数和遇到的错误。
// 如果 w 实现了 WriteString 方法，则优先使用该方法将 s 写入 w 中。
//否则，将 s 转换为 []byte，然后调用 w.Write 方法将数据写入 w 中。
// func WriteString(w Writer, s string) (n int, err error)

// ReadAtLeast 从 r 中读取数据到 buf 中，要求至少读取 min 个字节。返回读取的字节数和遇到的错误。
// func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
// 如果 min 超出了 buf 的容量，则 err 返回 io.ErrShortBuffer，否则：
// 1、读出的数据长度 == 0 ，则 err 返回 EOF。
// 读出的数据长度 < min，则 err 返回 io.ErrUnexpectedEOF。
// 读出的数据长度 >= min，则 err 返回 nil。

// ReadFull 的功能和 ReadAtLeast 一样，只不过 min = len(buf)
// func ReadFull(r Reader, buf []byte) (n int, err error)

func funcExample() {
	n, err := io.WriteString(os.Stdout, "Hello World!")
	fmt.Println(n, err)

	r := strings.NewReader("Hello World!")
	b := make([]byte, 15)

	n, err = io.ReadAtLeast(r, b, 20)
	fmt.Println(n, err, b[:n])

	r.Seek(0, 0)
	b = make([]byte, 15)
	n, err = io.ReadFull(r, b)
	fmt.Printf("%q   %d   %v\n", b[:n], n, err)

}

// LimitReader 对 r 进行封装，使其最多只能读取 n 个字节的数据。相当于对 r 做了一个切片 r[:n] 返回。底层实现是一个 *LimitedReader（只有一个 Read 方法）。
//func LimitReader(r Reader, n int64) Reader
func limitReaderExample() {
	r := strings.NewReader("Hello World")
	lr := io.LimitReader(r, 5)
	n, err := io.Copy(os.Stdout, lr)  // Hello
	fmt.Printf("\n%d   %v\n", n, err) // 5   <nil>
}

// MultiWriter MultiReader 将多个 Reader 封装成一个单独的 Reader，多个 Reader 会按顺序读取，当多个 Reader 都返回 EOF 之后，单独的 Reader 才返回 EOF，否则返回读取过程中遇到的任何错误。

func multiWriterExaminer() {
	r := strings.NewReader("Hello World")
	r.WriteTo(io.MultiWriter(os.Stdout, os.Stdout, os.Stdout))
}

// func MultiReader(readers ...Reader) Reader
func multiReaderExample() {
	r1 := strings.NewReader("Hello World")
	r2 := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r3 := strings.NewReader("abcdefghijklmnopqrstuvwxyz")

	b := make([]byte, 15)
	mr := io.MultiReader(r1, r2, r3)

	for n, err := 0, error(nil); err == nil; {
		n, err = mr.Read(b)
		fmt.Printf("%q\n", b[:n])
	}

	r1.Seek(0, 0)
	r2.Seek(0, 0)
	r3.Seek(0, 0)

	mr = io.MultiReader(r1, r2, r3)
	io.Copy(os.Stdout, mr)

}

func examples() {
	limitReaderExample()
	multiWriterExaminer()
	multiReaderExample()
}

// MultiWriter 将向自身写入的数据同步写入到所有 writers 中。
// func MultiWriter(writers ...Writer) Writer

// TeeReader 对 r 进行封装，使 r 在读取数据的同时，自动向 w 中写入数据。
//它是一个无缓冲的 Reader，所以对 w 的写入操作必须在 r 的 Read 操作结束之前完成。所有写入时遇到的错误都会被作为 Read 方法的 err 返回。
// func TeeReader(r Reader, w Writer) Reader

// CopyN 从 src 中复制 n 个字节的数据到 dst 中，返回复制的字节数和遇到的错误。只有当 written = n 时，err 才返回 nil。如果 dst 实现了 ReadFrom 方法，则优先调用该方法执行复制操作。
// Copy 从 src 中复制数据到 dst 中，直到所有数据都复制完毕，返回复制的字节数和遇到的错误。如果复制过程成功结束，则 err 返回 nil，而不是 EOF，因为 Copy 的定义为“直到所有数据都复制完毕”，所以不会将 EOF 视为错误返回。如果 src 实现了 WriteTo 方法，则调用 src.WriteTo(dst) 复制数据，否则如果 dst 实现了 ReadeFrom 方法，则调用 dst.ReadeFrom(src) 复制数据
//func Copy(dst Writer, src Reader) (written int64, err error)
// CopyBuffer 相当于 Copy，只不 Copy 在执行的过程中会创建一个临时的缓冲区来中转数据，而 CopyBuffer 则可以单独提供一个缓冲区，让多个复制操作共用同一个缓冲区，避免每次复制操作都创建新的缓冲区。如果 buf == nil，则 CopyBuffer 会自动创建缓冲区。
// func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)

func copyExample() {
	r := strings.NewReader("Hello World")
	buf := make([]byte, 32)

	n, err := io.CopyN(os.Stdout, r, 5)
	fmt.Printf("\n%d %v\n\n", n, err)

	r.Seek(0, 0)
	n, err = io.Copy(os.Stdout, r)
	fmt.Printf("\n%d, %v\n\n", n, err)

	r.Seek(0, 0)
	r2 := strings.NewReader("ABCDEFG")
	r3 := strings.NewReader("abcdefg")

	n, err = io.CopyBuffer(os.Stdout, r, buf)
	fmt.Printf("\n%d  %v", n, err)

	n, err = io.CopyBuffer(os.Stdout, r2, buf) // ABCDEFG
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>

	n, err = io.CopyBuffer(os.Stdout, r3, buf) // abcdefg
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>

}
