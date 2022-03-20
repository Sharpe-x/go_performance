package designa_pattern

import "fmt"

// 简单工厂模式
// go 语言没有构造函数一说 所以会定义NewXXX函数来初始化相关类
// NewXXX 函数返回接口时就是简单工厂模式

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI returns Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of the API implement
type hiAPI struct{}

// Say hi to name
func (hi *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI is one of the API implement
type helloAPI struct{}

// Say hello to name
func (hello *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
