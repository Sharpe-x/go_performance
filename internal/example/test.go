package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}
