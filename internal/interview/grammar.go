package main

import (
	"encoding/json"
	"fmt"
)

//  1、下面代码能运行吗？为什么。

type Param map[string]interface{}

type Show struct {
	Param
}

func test1() {
	s := new(Show) // new 不能初始化结构体中Param 的属性 所以报错
	s.Param["RMB"] = 1000
}

// 2、请说出下面代码存在什么问题。

type student struct {
	Name string
}

/*func zjl(v interface{}) {
	switch msg := v.(type) {
	case *student, student: // switch type 的 case T1，类型列表只有一个，那么 v := m.(type) 中的 v 的类型就是 T1 类型。
		// 如果是 case T1, T2，类型列表中有多个，那 v 的类型还是多对应接口的类型，也就是 m 的类型。
		// 所以这里 msg 的类型还是 interface{}，所以他没有 Name 这个字段，编译阶段就会报错。
		//fmt.Println(msg.Name)
	}
}*/

// 3、写出打印的结果。

type People struct {
	name string `json:"name"` //私有属性 name 也不应该加 json 的标签。
}

func test2() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
	// 小写开头的方法、属性或 struct 是私有的，同样，在 json 解码或转码的时候也无法上线私有属性的转换。
}

//4、下面的代码是有问题的，请说明原因。

type Person struct {
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("print: %v", p)
}

// 在 golang 中 String() string 方法实际上是实现了 String 的接口的，该接口定义在 fmt/print.go 中
func test3() {
	p := &Person{} // 在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。
	// 而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。
	p.String()
}

func main() {

}
