package main

import (
	"bytes"
	"fmt"
	"os"
)

//  bytes.NewBuffer 实现了很多基本的接口，可以通过 bytes 包学习接口的实现

func main() {
	buf := bytes.NewBuffer([]byte("hello world!"))
	b := make([]byte, buf.Len())
	n, err := buf.Read(b)
	fmt.Printf("%s   %v\n", b[:n], err) // hello world!   <nil>

	buf.WriteString("ABCDEFG\n")
	buf.WriteTo(os.Stdout) // ABCDEFG
	// buf is empty

	n, err = buf.Write(b)
	fmt.Printf("%d   %s   %v\n", n, buf.String(), err) // 12   hello World!   <nil>

	c, err := buf.ReadByte()
	fmt.Printf("%c   %s  %v\n", c, buf.String(), err) // h ello World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err) // e   llo World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err) // ello World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err) // e   llo World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err) // l lo World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err) // l o World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err) // lo World!   <nil>

}
