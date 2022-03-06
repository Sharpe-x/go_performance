package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name string
	Age  int32
}

func main() {
	// 可变参数是空接口类型时 传入空接口的切片时需要注意参数展开
	var a = []interface{}{1, 2, 3, "hello"}
	// 输出不同
	fmt.Println(a)    // [1 2 3 hello]
	fmt.Println(a...) // 1 2 3 hello

	// 数组是值传递
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 4
		fmt.Println(arr) // [4 2 3]
	}(x)
	fmt.Println(x) //[1 2 3]

	// 必要时使用切片
	y := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 4
		fmt.Println(arr) // [4 2 3]
	}(y)
	fmt.Println(y) // [4 2 3]

	students := []*Student{
		{
			Name: "one",
			Age:  1,
		},
		{
			Name: "Two",
			Age:  2,
		},
	}

	func(students []*Student) {
		students[0].Age = 100
		students[1].Name = "old people"
		fmt.Println(*students[0], *students[1])
	}(students)
	fmt.Println(*students[0], *students[1])

	//map 遍历时顺序不固定
	// map 是一种散列表实现 每次遍历的顺序可能不一样

	// 返回值被屏蔽
	/*	if Foo() != nil {
		fmt.Println("Foo failed")
	}*/

	//  Recover 必须在defer 函数中运
	fmt.Println(If(2 > 3, "hello", "hi").(string))

	//  := 表示声明并赋值，= 表示仅赋值。
	//变 量的作用域是大括号，因此在第一个 if 语句 if err == nil 内部重新声明且赋值了与外部变量同名的局部变量 err。对该局部变量的赋值不会影响到外部的 err。因此第二个 if 语句 if err
	//!= nil 不成立。所以只打印了 1 err。
	var err error
	if err == nil {
		err := fmt.Errorf("err") // 内部重新声明且赋值了与外部变量同名的局部变量 err。对该局部变量的赋值不会影响到外部的 err。
		fmt.Println(1, err)
	}
	if err != nil {
		fmt.Println(2, err)
	}

}

/*func Foo() (err error) {
	if err := Bar(); err != nil {
		return
	}
	return
}*/

func Bar() error {
	return errors.New("test")
}

// If Go 语言三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
