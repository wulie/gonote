package main

import "fmt"

// OrderMainDAO 定义接口
type OrderMainDAO interface {
	SaveOrderMainDAO()
}
type OrderDetailDAO interface {
	SaveOrderDetailDAO()
}

type OrderFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBOderMain struct {
}

func (r *RDBOderMain) SaveOrderMainDAO() {
	fmt.Println("save order main by rdb")
}

type RDBOrderDetail struct {
}

func (r *RDBOrderDetail) SaveOrderDetailDAO() {
	fmt.Println("save order detail by rdb")
}

type RDBOrderFactory struct {
}

func (r RDBOrderFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBOderMain{}
}

func (r RDBOrderFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBOrderDetail{}
}

type XMLOrderMain struct {
}

func (x *XMLOrderMain) SaveOrderMainDAO() {
	fmt.Println("save order main by xml")
}

type XMLOrderDetail struct {
}

func (x *XMLOrderDetail) SaveOrderDetailDAO() {
	fmt.Println("save order detail by xml")
}

type XMLOrderFactory struct {
}

func (r XMLOrderFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLOrderMain{}
}

func (r XMLOrderFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLOrderDetail{}
}

func main() {
	var fac OrderFactory
	fac = &RDBOrderFactory{}
	getOrder(fac)
	fac = &XMLOrderFactory{}
	getOrder(fac)
}

func getOrder(fac OrderFactory) {
	fac.CreateOrderMainDAO().SaveOrderMainDAO()
	fac.CreateOrderDetailDAO().SaveOrderDetailDAO()
}
