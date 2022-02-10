package main

import (
	"bytes"
	"fmt"
)

func main() {
	var ch string
	fmt.Scanf("%s", &ch)

	buffer := new(bytes.Buffer)
	_, err := buffer.Write([]byte(ch))
	if err == nil {
		fmt.Println("写入第一个字节成功！开始读取改字节")
		r, n, err := buffer.ReadRune()
		fmt.Println(string(r), n, err)
	} else {
		fmt.Println("写入错误")
	}

}
