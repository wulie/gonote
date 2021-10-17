package main

import "fmt"

type MotherBoard struct {
}

func (m *MotherBoard) Start() {
	fmt.Println("system is starting....")
}

func (m *MotherBoard) Reboot() {
	fmt.Println("system is reboot....")
}

type Command interface {
	Execute()
}

type StartCommand struct {
	md *MotherBoard
}

func NewStartCommand(md *MotherBoard) *StartCommand {
	return &StartCommand{md: md}
}
func (s *StartCommand) Execute() {
	s.md.Start()
}

type RebootCommand struct {
	md *MotherBoard
}

func NewRebootCommand(md *MotherBoard) *RebootCommand {
	return &RebootCommand{md: md}
}

func (s *RebootCommand) Execute() {
	s.md.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}
func (b *Box) PressButton2() {
	b.button2.Execute()
}
func main() {
	md := &MotherBoard{}
	startCommand := NewStartCommand(md)
	rebootCommand := NewRebootCommand(md)

	box := NewBox(startCommand, rebootCommand)
	box.PressButton1()
	box.PressButton2()
	box = NewBox(rebootCommand, startCommand)
	box.PressButton1()
	box.PressButton2()
}
