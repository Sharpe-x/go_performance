package main

import "fmt"

const (
	a = iota
	b = iota
	f = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(c)
	fmt.Println(d)

}
