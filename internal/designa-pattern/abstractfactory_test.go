package designa_pattern

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func TestRdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)

	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
