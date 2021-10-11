package main

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecReq() string
}

type ApapteeImpl struct {
}

func (a *ApapteeImpl) SpecReq() string {
	return "adapter method"
}

func NewAdaptee() Adaptee {
	return &ApapteeImpl{}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecReq()
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

func main() {
	adp := NewAdaptee()
	newAdapter := NewAdapter(adp)
	println(newAdapter.Request())
}

//一个接口是 暴露的接口ia ，另外一个接口才是真正要调用的ib ,
//定义一个结构体b，实现真正调用的接口ib，实现一个b的工程方法
//定义新的结构体a,包含ib， 实现接口ia，实现一个a的工程方法，参数是ib
//使用方法
//1.创建b
//2.创建啊a 参数是b
//3.调用ia的方法实际是调用ib的方法
