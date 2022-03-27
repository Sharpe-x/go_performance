package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numStr := scanner.Text()
		n, err := strconv.Atoi(numStr)
		if err != nil || n == 0 {
			continue
		}
		//fmt.Println(n / 2)
		fmt.Println(getAns(n))
	}
}

func getAns(n int) int {
	ans := 0
	for n >= 3 {
		ans += n / 3
		n = n/3 + n%3
	}

	if n == 2 {
		ans += 1
	}
	return ans
}
