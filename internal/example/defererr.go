package main

import (
	"fmt"
	"time"
)

func test() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("err is not nil = ", err.Error())
			err = nil
		}
	}()

	return crateErr()
}

func crateErr() error {
	return fmt.Errorf("create err %d", time.Now().UnixNano())
}

func main() {
	err := test()
	fmt.Println(err)
}
