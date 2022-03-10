package main

import "fmt"

func CreateCase() {
	// 1.Go 语言的数组有2 种创建方式
	// 第一种 ： 显示指定数组大小
	arr1 := [2]int{1, 2}
	// 第二种：使用【...】T 声明数组 Go语言会在编译期间通过源代码推导数组大小
	arr2 := [...]int{1, 2, 3, 4, 5}

	// 不考虑逃逸分析的话 对由于字面量组成的数组 根据数组元素数量的大小 编译器会做出优化
	// 当元素少于或者等于4 个时 会直接将元素放在栈上
	// 当元素大于4个时 会将数组中的元素放到静态区并在运行时取出 （静态区初始化然后复制到栈）

	fmt.Println(arr1)
	fmt.Println(arr2)

	changeArr(arr1)
	fmt.Println(arr1)

	changeArrPtr(&arr1)
	fmt.Println(arr1)

}

// 数组传递时值传递 不会改变外边的值
func changeArr(arr [2]int) {
	arr[0] = arr[0] * 10
	arr[1] = arr[1] * 100
	fmt.Println("changeArr inner arr", arr)
}

// 指针传递的话 会改变外部的值
func changeArrPtr(arr *[2]int) {
	arr[0] = arr[0] * 10
	arr[1] = arr[1] * 100
	fmt.Println("changeArrPtr inner arr", arr)
}

func main() {
	CreateCase()
}

func changeNumber(i int) {
	i = i * i
}
