package main

// compare to a c++ example: http://www.tutorialspoint.com/cplusplus/cpp_interfaces.htm

import (
	"fmt"
)

// interface
type Shape interface {
	Area() float64
	GetWidth() float64
	GetHeight() float64
	SetWidth(float64)
	SetHeight(float64)
}

// reusable part, only implement SetWidth and SetHeight method of the interface
// {

type WidthHeight struct {
	width  float64
	height float64
}

func (this *WidthHeight) SetWidth(w float64) {
	this.width = w
}
func (this *WidthHeight) SetHeight(h float64) {
	this.height = h
}
func (this *WidthHeight) GetWidth() float64 {
	// fmt.Println("in WidthHeight.GetWidth")
	return this.width
}
func (this *WidthHeight) GetHeight() float64 {
	fmt.Println("in WidthHeight.GetHeight")
	return this.height
}

// }

type Rectangle struct {
	WidthHeight
}

func (this *Rectangle) Area() float64 {
	fmt.Printf("%T\n", this)
	return this.GetWidth() * this.GetHeight() / 2
}

// override
func (this *Rectangle) GetHeight() float64 {
	fmt.Println("in Rectangle.GetHeight")
	// in case you still needs the WidthHeight's GetHeight method
	return this.WidthHeight.GetHeight()
}

func main() {
	var r Rectangle
	var i Shape = &r
	i.SetWidth(4)
	i.SetHeight(6)

	fmt.Printf("%v %T", i, i)
	fmt.Println("width: ", i.GetWidth())
	fmt.Println("height: ", i.GetHeight())
	fmt.Println("area: ", i.Area())

}
