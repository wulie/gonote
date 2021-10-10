package main

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Direct struct {
	builder Builder
}

func NewDirect(builder Builder) *Direct {
	return &Direct{builder: builder}
}

func (d *Direct) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}
func (b *Builder2) GetResult() int {
	return b.result
}
func main() {
	builder1 := &Builder1{}
	direct := NewDirect(builder1)
	direct.Construct()
	println(builder1.GetResult())

	builder2 := &Builder2{}
	NewDirect(builder2).Construct()
	println(builder2.GetResult())

}
