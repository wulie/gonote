package main

import "fmt"

type Operator interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}
type PlusOperatorFactory struct {
}

func (p PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}

type MinusOperator struct {
	*OperatorBase
}

type MinusOperatorFactory struct {
}

func (m MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}

func (m *MinusOperator) Result() int {
	return m.a - m.b
}

type MultiOperator struct {
	*OperatorBase
}

func (m *MultiOperator) Result() int {
	return m.a * m.b
}

type MultiOperatorFactory struct {
}

func (m MultiOperatorFactory) Create() Operator {
	return &MultiOperator{
		&OperatorBase{},
	}
}

func main() {
	var fac OperatorFactory

	fac = PlusOperatorFactory{}
	i := compute(fac, 3, 5)
	fmt.Println(i)
	fac = MinusOperatorFactory{}
	i = compute(fac, 3, 5)
	fmt.Println(i)
	fac = MultiOperatorFactory{}
	i = compute(fac, 3, 5)
	fmt.Println(i)
}

func compute(fac OperatorFactory, a, b int) int {
	create := fac.Create()
	create.SetA(a)
	create.SetB(b)
	result := create.Result()
	return result
}
