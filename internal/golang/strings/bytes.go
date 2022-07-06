package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	bytes := []byte("")
	fmt.Println(bytes == nil)
	fmt.Println(strings.Split("A;", ";"))
	fmt.Println(len(strings.Split("A;", ";")))
	fmt.Println()
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixMicro())
	fmt.Println(now.UnixNano())
}
