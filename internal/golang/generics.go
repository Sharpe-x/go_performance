package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Min[F Number](x, y F) F {
	if x > y {
		return y
	}
	return x
}

func GMin[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func Scale[E constraints.Integer](s []E, c E) []E {

	retSlice := make([]E, len(s))
	for i, e := range s {
		retSlice[i] = e * c
	}
	return retSlice
}

func Scale2[S ~[]E, E constraints.Integer](s S, e E) S {
	retSlice := make(S, len(s))

	for i, e2 := range s {
		retSlice[i] = e2 * e
	}

	return retSlice
}

type Pointer []int32

func (p Pointer) String() string {
	var sum int32
	for _, num := range p {
		sum += num
	}
	return strconv.Itoa(int(sum))
}

func ScalePointer(p Pointer) {
	fmt.Println(Scale(p, 2))
}

func ScalePointer2(p Pointer) {
	fmt.Println(Scale2(p, 2))
}

type Slice[T int | float64 | float32 | int64] []T

type MyMap[KEY int | string, VALUE float64 | string] map[KEY]VALUE

// type NewType[T *int,] []T 错误写法

// NewType 单个类型约束 可以加逗号消除歧义
type NewType[T *int, ] []T

// NewType2 多个类型约束 使用interface 包裹
type NewType2[T interface{ *int | *int64 }] []T

// 泛型类型的套娃

type MySlice[T constraints.Integer | constraints.Float] []T

func (ms MySlice[T]) Sum() T {
	var sum T
	for _, t := range ms {
		sum += t
	}
	return sum
}

type MyIntSlice[T int | int32] MySlice[T]

type MySliceMap[T MySlice[float32] | MySlice[float64]] map[string]T

type MySliceMap2[T float32 | float64] map[string]MySlice[T]

func main() {
	/*	fmt.Println(Min(1, 2))
		fmt.Println(Min(12345.3932, 12345.3934))
		fmt.Println(GMin(1, 2))*/

	//a := []int32{1, 2, 3, 4, 5}
	//fmt.Println(Scale(a, 10))

	//ScalePointer(a)
	//ScalePointer2(a)

	var intSlice Slice[int] = []int{1, 2, 3}
	fmt.Printf("intSlice name %T\n", intSlice)

	float32Slice := Slice[float32]([]float32{1.0, 2.0, 3.0})
	fmt.Printf("float32Slice name %T\n", float32Slice)

	var myMap MyMap[int, string] = map[int]string{
		1: "一",
		2: "二",
	}

	fmt.Printf("myMap name %T\n", myMap)
	fmt.Println(myMap)

	var a MySlice[int] = []int{1, 2, 3, 4, 5}
	fmt.Println(a.Sum())

	var a2 MySlice[float64] = []float64{1.0, 2.1, 3.2}
	fmt.Println(a2.Sum())

}
