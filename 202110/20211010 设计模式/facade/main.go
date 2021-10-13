package main

type AModule interface {
	Test() string
}

type aModule struct {
}

func NewAModule() AModule {
	return &aModule{}
}

func (a *aModule) Test() string {
	return "module a is running"
}

type BModule interface {
	Test() string
}

type bModule struct {
}

func NewBModule() BModule {
	return &bModule{}
}

func (b *bModule) Test() string {
	return "module b is running"
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModule
	b BModule
}

func NewAPIModule() API {
	return &apiImpl{
		a: NewAModule(),
		b: NewBModule(),
	}
}

func (a *apiImpl) Test() string {
	test := a.a.Test()
	s := a.b.Test()
	return test + "\n" + s
}

func main() {
	api := NewAPIModule()
	println(api.Test())
}

//一个系统里面有很多子系统，他们的接口都相同， 并且很多业务也都是在一起，可以封装抽象一个接口，供用户调用， 这样可以减少用户操作对象的数目
//模块A接口  有TEST方法  用a实现
//模块B接口 有TEST方法 用b实现
//抽象一个API接口， 包含 A和B，并且也定义一个TESt 方法 用apiImpl 实现
//这样就可以用过创建apiImpl 对象去操作接口A和B， 减少单独调用使用a和b
