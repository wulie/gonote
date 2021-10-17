package main

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
		context:   "",
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(subject *Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) Update(subject *Subject) {
	fmt.Printf("%s received message: %s\n", r.name, subject.context)
}

func main() {
	wxh := NewReader("wxh")
	jsc := NewReader("jsc")
	subject := NewSubject()
	subject.Attach(wxh)
	subject.Attach(jsc)
	subject.UpdateContext("明天上班")
}

//观察者模式用于触发联动。
//
//一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。
