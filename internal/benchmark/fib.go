package benchmark

import (
	"math/rand"
	"time"
)

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}

func GenerateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}

	return nums
}

func Generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
