package designa_pattern

import "fmt"

// 抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的
// 如果抽象工厂退化生成的对象无关联则成为工厂函数模式
// 比如本例子中使用RDB和XML存储对的信息，抽象工厂分别能生成相关的主订单信息和订单详情信息
// 如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换不同的存储方式了。

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象工厂模式接口
type DAOFactory interface {
	CreateOrderMainDao() OrderMainDAO
	CreateOrderDetailDao() OrderDetailDAO
}

// RDBMainDAO 为关系型数据库的OrderMainDAO的实现
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO 为关系型数据库的OrderDetailDAO的实现
type RDBDetailDAO struct{}

// SaveOrderDetail ...
func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory 是RDB抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &RDBMainDAO{}
}
func (*RDBDAOFactory) CreateOrderDetailDao() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML 存储
type XMLMainDAO struct{}

// SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// XMLDetailDAO XML 存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

// XMLDAOFactory XML 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderDetailDao() OrderDetailDAO {
	return &XMLDetailDAO{}
}

func (*XMLDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &XMLMainDAO{}
}
