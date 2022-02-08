package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(strings.Count(str, ""))
	fmt.Println(str[0 : 0+1])
	fmt.Println(strings.Count(str, ""))
	fmt.Println(str[1 : 1+1])
	fmt.Println(strings.Count(str, ""))
}
