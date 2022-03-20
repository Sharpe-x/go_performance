package designa_pattern

// 工厂方法模式使用子类的方式延迟生成对象到子类中实现
// Go 中不存在继承 所以使用匿名组合来实现

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类 封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA sets the operator A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB sets the operator B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory 是PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

// MinusOperatorFactory 是MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
