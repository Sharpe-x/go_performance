package streamexample

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
)

// NewEncoder base64 编码成字符串 需要一个io.Writer作为输出目标，并用返回的WriteCloser的Write方法将结果写入目标
// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
// encoderExample base64编码成字符串
func encoderExample() {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)

	buffer := new(bytes.Buffer)
	encoder2 := base64.NewEncoder(base64.StdEncoding, buffer)
	encoder2.Write(input)
	fmt.Println()
	fmt.Println(buffer.String())
}

// []byte和struct之间正反序列化 这种场景经常用在基于字节的协议上，比如有一个具有固定长度的结构：

type Protocol struct {
	Version  uint8
	BodyLen  uint16
	Reserved [2]byte
	Unit     uint8
	Value    uint32
}

// 通过一个[]byte 来反序列化得到这个protocol 一种思路是遍历这个[]byte 然后逐一赋值
// func Read(r io.Reader, order ByteOrder, data interface{}) error encoding/binary包中有个方便的方法
// 从一个io.Reader中读取字节，并已order指定的端模式，来给填充data（data需要是fixed-sized的结构或者类型）。要用到这个方法首先要有一个io.Reader
// var p Protocol
//var bin []byte
// binary.Read(bytes.NewReader(bin), binary.LittleEndian, &p)
// 换句话说，我们将一个[]byte转成了一个io.Reader。
// 需要将Protocol序列化得到[]byte，使用encoding/binary包中有个对应的Write方法： func Write(w io.Writer, order ByteOrder, data interface{}) error
// var p Protocol
//buffer := new(bytes.Buffer)
////...
//binary.Writer(buffer, binary.LittleEndian, p)
//bin := buffer.Bytes()

// 从流中按行读取
// 比如对于常见的基于文本行的HTTP协议的读取，我们需要将一个流按照行来读取。本质上，我们需要一个基于缓冲的读写机制（读一些到缓冲，然后遍历缓冲中我们关心的字节或字符）。在Go中有一个bufio的包可以实现带缓冲的读写：
// unc NewReader(rd io.Reader) *Reader
//func (b *Reader) ReadString(delim byte) (string, error)

// 这个ReadString方法从io.Reader中读取字符串，直到delim，就返回delim和之前的字符串。如果将delim设置为\n，相当于按行来读取了：
// var conn net.Conn
////...
//reader := NewReader(conn)
//for {
//    line, err := reader.ReadString([]byte('\n'))
//    //...
