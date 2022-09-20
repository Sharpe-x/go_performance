package main

import "fmt"

// Go中字符串是只读的

type demo struct {
	hello string
}

func (d *demo) Clear() {
	fmt.Println(d.hello)
}

func main() {
	item := new(demo)
	defer print(item)
	defer item.Clear()
	*item = demo{
		hello: "hi",
	}
}

func (d *demo) print() {
	fmt.Println(d)
}
