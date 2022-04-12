package main

import (
	"fmt"
	"net/url"
)

type temp struct {
	Id string
}

type tempList struct {
	count int64
	list  []temp
}

func main() {
	fmt.Println(url.QueryEscape("12345/67890"))
	str := "1234567890"
	fmt.Println(url.QueryEscape(str))
	m, _ := url.QueryUnescape(str)
	fmt.Println(m)

	t := tempList{
		count: 1,
	}
	t.list = append(t.list, temp{
		Id: "zero",
	})

	fmt.Println(t)
}
