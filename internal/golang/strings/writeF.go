package main

import (
	"bufio"
	"fmt"
	"os"
)

var test = []string{
	"1qaz",
	"2wsx",
	"3edc",
}

func main() {
	fileName := "result.txt"

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, msg := range test {
		_, _ = fmt.Fprintln(w, msg)
	}

	_ = w.Flush()
}