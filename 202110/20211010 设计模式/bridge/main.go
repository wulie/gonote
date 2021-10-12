package main

import "fmt"

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImpler interface {
	Send(text, to string)
}

type MessageSMS struct {
}

func viaSMS() MessageImpler {
	return &MessageSMS{}
}
func (v *MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct {
}

func viaEmail() MessageImpler {
	return &MessageEmail{}
}
func (m *MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

type CommonMessage struct {
	method MessageImpler
}

func NewCommonMessage(method MessageImpler) AbstractMessage {
	return &CommonMessage{method: method}
}
func (c *CommonMessage) SendMessage(text, to string) {
	c.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImpler
}

func NewUrgencyMessage(method MessageImpler) AbstractMessage {
	return &UrgencyMessage{method: method}
}

func (u *UrgencyMessage) SendMessage(text, to string) {
	u.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

func main() {
	NewCommonMessage(viaSMS()).SendMessage("have a drink?", "Bob")
	fmt.Println()
	NewCommonMessage(viaEmail()).SendMessage("have a drink?", "Bob")
	fmt.Println()
	NewUrgencyMessage(viaSMS()).SendMessage("have a drink?", "Bob")
	fmt.Println()
	NewUrgencyMessage(viaEmail()).SendMessage("have a drink?", "Bob")
}

//有多个变化维度（多个变化的原因）的系统的时候，可以使用桥接模式
//如例子上面， 发送信息的方式有SMS短信和邮件Email
//急迫程度也有两种 普通语气Common和急切语气Urgency
// 先定义抽象发送信息的接口ia
// 在定义具体发消息的方式接口ib
//定义实现具体发消息接口的结构体SMS和Email 实现ib
//实现对象创建方法
//定义语气发消息的接口体 Common 和Urgency ， 包含ib 属性
//实现对象 创建方法
//使用方法-----------------------------------------------
// 组合使用即可发aa
