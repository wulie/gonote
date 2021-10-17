package main

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "sound,image"
	fmt.Printf("CDDriver: reading data :%s\n", c.Data)
	GetMediatorInstance().change(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]
	fmt.Printf("CPU: split data wuth sound:%s,Sound %s", c.Sound, c.Video)
	GetMediatorInstance().change(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard display: %s", v.Data)
	GetMediatorInstance().change(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard play:%s", s.Data)
	GetMediatorInstance().change(s)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) change(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)

	}
}

func main() {
	mediator := GetMediatorInstance()
	mediator.CPU = &CPU{}
	mediator.CD = &CDDriver{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}
	mediator.CD.ReadData()
	fmt.Println(mediator.CD.Data)
	println(mediator.CPU.Sound)
	println(mediator.CPU.Video)

}

//中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。
//
//例子中的中介者使用单例模式生成中介者。
//
//中介者的change使用switch判断类型。
