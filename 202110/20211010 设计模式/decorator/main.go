package main

import "fmt"

type Component interface {
	Calc() int
}

type ConCreateComponent struct {
}

func (c *ConCreateComponent) Calc() int {

	return 0
}

type AddComponent struct {
	Component
	num int
}

func (a *AddComponent) Calc() int {
	return a.Component.Calc() + a.num
}

type MulComponent struct {
	Component
	num int
}

func WarpMulComponent(c Component, num int) Component {
	return &MulComponent{
		Component: c,
		num:       num,
	}
}
func (a *MulComponent) Calc() int {
	return a.Component.Calc() * a.num
}

func WarpAddComponent(c Component, num int) Component {
	return &AddComponent{
		Component: c,
		num:       num,
	}
}

func main() {
	var calc Component = &ConCreateComponent{}
	println(calc.Calc())
	calc = WarpMulComponent(calc, 10)
	println(calc.Calc())
	fmt.Println(calc)
	calc = WarpAddComponent(calc, 8)
	println(calc.Calc())
	fmt.Println(calc)

}

//装饰模式装饰模式使用对象组合的方式动态改变或增加对象行为， 顺序不同结果不同， 不同于工厂模式
