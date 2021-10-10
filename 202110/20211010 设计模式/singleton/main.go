package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	foo()
}

type singleton struct {
}

func (s *singleton) foo() {
	fmt.Println("foo")
}

//单例模式变量
var (
	instance *singleton
	once     sync.Once
)

func getInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	a := getInstance()
	b := getInstance()
	if a == b {
		fmt.Println("a=b ")
	}
}
