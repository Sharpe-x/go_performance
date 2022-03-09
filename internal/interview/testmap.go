package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := make(map[string]int)

	fillMap(test)

	fmt.Println(test)

}

func fillMap(val map[string]int) {
	for i := 0; i < 10; i++ {
		val["fill"+"_"+strconv.Itoa(i)] = 1
	}
}
