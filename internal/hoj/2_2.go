package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	scanner.Scan()
	substr := scanner.Text()
	fmt.Println(strings.Count(strings.ToUpper(str), strings.ToUpper(substr)))
}
