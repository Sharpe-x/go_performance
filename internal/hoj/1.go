package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _, err := inputReader.ReadLine()
	if err != nil {
		panic(err)
	}
	str := strings.TrimSpace(string(input))
	slice := strings.Split(str, " ")
	fmt.Println(len(slice[len(slice)-1]))
}
