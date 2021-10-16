package main

type Subject interface {
	Do() string
}

type RealObject struct {
}

func (receiver *RealObject) Do() string {
	return "real"
}

type Proxy struct {
	real RealObject
}

func (p *Proxy) Do() string {

	pre := "before:"
	do := p.real.Do()
	pre += do
	pre += ":after"
	return pre
}

func main() {

	var sub Subject
	sub = &Proxy{}
	println(sub.Do())

}

//代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理虚代理
//COW代理
//远程代理
//保护代理
//Cache 代理
//防火墙代理
//同步代理
//智能指引
//等。。。。
