package concurrency

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	for i := 0; i < 10000; i++ {
		Order()
		fmt.Println()
	}

}
