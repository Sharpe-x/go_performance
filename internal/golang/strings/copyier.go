package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type notifyConfig struct {
	RobotUrl string
	Mention  []string
	JumpUrl  string
}

func notify() {
	config := notifyConfig{
		RobotUrl: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=a0245510-7ce0-4ce1-9458-2612a931ff81",
		Mention:  []string{"sharpezhang", "charcliu"},
		JumpUrl:  "https://dev.ops.tsign.woa.com/tool-brand-cues",
	}

	bytes, _ := json.Marshal(&config)
	fmt.Println(string(bytes))
}

func main() {
	notify()
	learnReflect()
	learnMore()
	learnMore2()
	refactorStruct()
}

func learnReflect() {
	var x = 3.49
	v := reflect.ValueOf(x)
	fmt.Println(v.CanSet())
	vPtr := reflect.ValueOf(&x)
	fmt.Println(vPtr.CanSet())
	fmt.Println(vPtr.Elem().CanSet())
	if vPtr.Elem().CanSet() {
		vPtr.Elem().SetFloat(789.7654)
	}
	fmt.Println(x)
}

type Foo interface {
	Name() string
}

type FooStruct struct {
	A string
	b int
}

func (f FooStruct) Name() string {
	return f.A
}

type FooPointer struct {
	A string
}

func (f *FooPointer) Name() string {
	return f.A
}

func learnMore() {
	{
		a := []int{1, 2, 3}
		val := reflect.ValueOf(&a)
		val.Elem().SetLen(2)
		val.Elem().Index(0).SetInt(4)
		fmt.Println(a)
	}

	{
		a := map[int]string{
			1: "foo1",
			2: "foo2",
		}
		val := reflect.ValueOf(a)
		key3 := reflect.ValueOf(3)
		val3 := reflect.ValueOf("foo3")
		val.SetMapIndex(key3, val3)
		fmt.Println(val)
	}

	{
		a := map[int]string{
			1: "foo1",
			2: "foo2",
		}
		val := reflect.ValueOf(&a)
		key3 := reflect.ValueOf(3)
		val3 := reflect.ValueOf("foo3")
		val.Elem().SetMapIndex(key3, val3)
		fmt.Println(val)
	}

	{
		a := FooStruct{}
		val := reflect.ValueOf(&a)
		val.Elem().FieldByName("A").SetString("foo2")
		fmt.Println(a)
	}
	{
		a := &FooPointer{}
		val := reflect.ValueOf(a)
		val.Elem().FieldByName("A").SetString("foo2")
		fmt.Println(a)
	}
	{
		a := FooStruct{}
		val := reflect.ValueOf(&a)
		fmt.Println(val.Elem().FieldByName("b").CanSet())
		fmt.Println(val.Elem().FieldByName("b").CanAddr())
	}
}

func learnMore2() {
	var i *int
	v := reflect.ValueOf(i)
	v2 := v.Elem()

	fmt.Println(v2.IsValid())
	fmt.Println(reflect.ValueOf(nil).IsValid())
	fmt.Println(reflect.Indirect(v).IsValid())
}

type Turbo struct {
	Name string
	Age  int
	notifyConfig
}

func refactorStruct() {
	var turbo = &Turbo{
		Name: "sharpe",
		Age:  18,
		notifyConfig: notifyConfig{
			RobotUrl: "http://localhost",
			JumpUrl:  "http://localhost:80",
			Mention:  []string{"1", "2"},
		},
	}

	types := reflect.TypeOf(turbo)

	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}

	for i := 0; i < types.NumField(); i++ {
		tf := types.Field(i)
		fmt.Printf("字段名称:%v,字段类型:%v\n", tf.Name, tf.Type)
		fmt.Printf("字段名称:%v是不是匿名字段:%v\n", tf.Name, tf.Anonymous)
	}
}
