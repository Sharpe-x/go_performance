package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	file, err := os.Open("flowId.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufferReader := bufio.NewReader(file)

	for {
		a, _, c := bufferReader.ReadLine()
		if c == io.EOF {
			break
		}

		fmt.Println(string(a))
	}

	for i := 0; i < 100; i++ {
		index := i
		go func() {
			fmt.Println(index)
		}()
	}

	time.Sleep(time.Second)
}
