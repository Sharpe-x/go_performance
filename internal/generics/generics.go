package main

import (
	"fmt"
)

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

/*func SumNumbers[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Number 约束
type Number interface {
	int64 | float64
}

func SumNumberByInterface[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}*/

func main() {
	ints := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	floats := map[string]float64{
		"one":   1.1,
		"two":   2.2,
		"three": 3.3,
	}

	fmt.Println(SumInts(ints), SumFloats(floats))
	/*	fmt.Println(SumNumbers[string, int64](ints))

		fmt.Println(SumNumbers(floats))

		fmt.Println(SumNumberByInterface(ints))*/

}
