package main

import "fmt"

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data is %s", filename)
	return &ImageFlyweight{data: data}
}

func (f *ImageFlyweight) Data() string {
	return f.data
}

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFlyweightFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFlyweightFactory == nil {
		imageFlyweightFactory = &ImageFlyweightFactory{maps: make(map[string]*ImageFlyweight)}
	}
	return imageFlyweightFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		//不用用 := 会导致image作用域发生变化
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{image}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}

func main() {

	viewer := NewImageViewer("image.png")
	viewer.Display()

}

//元享模式，享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元，使多个对象共享，从而节省内存以及减少对象数量。
