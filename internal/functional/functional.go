package main

import (
	"fmt"
	"time"
)

func multiply(a, b int) int {
	return a * b
}

// MultiplyFunc 在不修改现有 multiply 函数代码的前提下计算乘法运算的执行时间，显然，这可以引入装饰器模式来实现。
// 通过 type 语句为匿名函数类型设置了别名 MultiPlyFunc
//  续就可以用这个类型别名来声明对应的函数类型参数和返回值，提高代码可读性。
type MultiplyFunc func(int, int) int

// 装饰器模式实现代码 execTime 函数 以 MultiPlyFunc 类型为参数和返回值 的高阶函数
// 所谓高阶函数，就是接收其他函数作为参数传入，或者把其他函数作为结果返回的函数。
func execTime(f MultiplyFunc) MultiplyFunc {
	return func(a int, b int) int {
		// 在返回的 MultiPlyFunc 类型匿名函数体中，真正执行乘法运算函数 f 前，
		//先通过 time.Now() 获取当前系统时间，并将其赋值给 start 变量；
		start := time.Now()
		c := f(a, b)                            //执行 f 函数，将返回值赋值给变量 c
		end := time.Since(start)                // time.Since(start) 计算从 start 到现在经过的时间，也就是 f 函数执行耗时，将结果赋值给 end 变量并打印出来；
		fmt.Println("--- 执行耗时: ", end, " ----") // 最后返回 f 函数运行结果 c 作为最终返回值。
		return c
	}
}

// 乘法运算函数（位运算）
func multiply2(a, b int) int {
	return a << b
}

func main() {
	a := 2
	b := 9
	//c := multiply(a, b)
	fmt.Println("算术运算：")
	decorator := execTime(multiply)
	c := decorator(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, c)

	fmt.Println("位运算：")
	decorator = execTime(multiply2)
	c = decorator(a, b)
	fmt.Printf("%d << %d = %d\n", a, b, c)

	testMapForEach()
	Test1()
	fmt.Println()
	fmt.Println(calculator(1, 111, "+"))

	fmt.Println(factorial(10))
	fmt.Println(factorialTailRecursive(10))
	testFib()
}

func mapForEach(arr []string, fn func(it string) int) []int {
	var newArray []int
	for _, s := range arr {
		newArray = append(newArray, fn(s))
	}
	return newArray
}

func testMapForEach() {
	var arr = []string{"hello", "handler", "execTime"}

	var out = mapForEach(arr, func(str string) int {
		return len(str)
	})
	fmt.Println(out)

}

// 什么是 Functional Programming
// 首先我们需要研究一下什么是高阶函数编程？所谓的 Functional Programming，一般被译作函数式编程（以 λ演算1 为根基）。
// 函数式编程，是指忽略（通常是不允许）可变数据（以避免它处可改变的数据引发的边际效应），
//忽略程序执行状态（不允许隐式的、隐藏的、不可见的状态），通过函数作为入参，函数作为返回值的方式进行计算
//，通过不断的推进（迭代、递归）这种计算，从而从输入得到输出的编程范式
// 在函数式编程范式中，没有过程式编程所常见的概念：语句，过程控制（条件，循环等等）。此外，
//在函数式编程范式中，具有引用透明（Referential Transparency）的特性，此概念的含义是函数的运行仅仅和入参有关，入参相同则出参必然总是相同，函数本身（被视作f(x)）所完成的变换是确定的。

// 总结一下，函数式编程具有以下的表征：
// No Data mutations 没有数据易变性
//No implicit state 没有隐式状态
//No side effects 没有边际效应（没有副作用）
//Pure functions only 只有纯粹的函数，没有过程控制或者语句
//First-class function 头等函数身份
//First-class citizen 函数具有一等公民身份
//Higher-order functions 高阶函数，可以出现在任何地方
//Closures 闭包 - 具有上级环境捕俘能力的函数实例
//Currying 柯里化演算2 - 规约多个入参到单个，等等
//Recursion 递归运算 - 函数嵌套迭代以求值，没有过程控制的概念
//Lazy evaluations / Evaluation strategy 惰性求值 - 延迟被捕俘变量的求值到使用时
//Referential transparency 引用透明性 - 对于相同的输入，表达式的值必须相同，并且其评估必须没有副作用

// 在 Golang 中，高阶函数很多时候是为了实现某种算法的关键粘合剂。
// 基本的闭包结构
//递归
//函子/运算子
//惰性计算
//可变参数：Functional Options

// 基本的闭包（Closure）结构Permalink

// 在函数、高阶函数身属一阶公民的编程语言中，你当然可以将函数赋值为一个变量、复制给一个成员，作为另一函数的参数（或之一）进行传参，作为另一函数的返回值（或之一）。

type Handler func(a int)

func xc(pa int, handler Handler) {
	handler(pa)
}

func Test1() {
	xc(1, func(a int) { // <- 再写一遍原型吧
		print(a)
	})
}

// 算子通常是一个简单函数（但也未必如此），总控部分通过替换不同算子来达到替换业务逻辑的实际实现算法：
func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }

var operators map[string]func(a, b int) int

func init() {
	operators = map[string]func(a int, b int) int{
		"+": add,
		"-": sub,
	}
}

func calculator(a, b int, op string) int {
	if fn, ok := operators[op]; ok && fn != nil {
		return fn(a, b)
	}
	return 0
}

// 递归 RecursionPermalink
// 斐波拉契，阶乘，Hanoi 塔，分形等是典型的递归问题。

// 阶乘法
func factorial(num int) int {
	result := 1
	for ; num > 0; num-- {
		result *= num
	}
	return result
}

// Functional Programming 的风格重新实现阶乘法
func factorialTailRecursive(num int) int {
	return factorialRecursive(1, num)
}
func factorialRecursive(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}
	return factorialRecursive(accumulator*val, val-1)
}

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

// 采用高阶函数的递归
func testFib() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
}

// Functional OptionsPermalink
// 旧的代码

type Holder struct {
	a int
	b bool
}

func New(a int) *Holder {
	return &Holder{
		a: a,
	}
}

// 后来加了一个布尔值b 于是修改 New

type Holder2 struct {
	a int
	b bool
}

func NewHolder2(a int, b bool) *Holder2 {
	return &Holder2{
		a: a,
		b: b,
	}
}

// 后来又加了一个 。。。 再修改
// 且不说改来改去烦不烦 使用者就要mmp 了

// 所以新的方式

type Opt func(holder *Holder)

func NewHolder(opts ...Opt) *Holder {
	h := &Holder{a: -1}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithInt(a int) Opt {
	return func(holder *Holder) {
		holder.a = a
	}
}

func WithBool(b bool) Opt {
	return func(holder *Holder) {
		holder.b = b
	}
}

// 不会影响用户使用
var holder = NewHolder(WithInt(1), WithBool(true))
