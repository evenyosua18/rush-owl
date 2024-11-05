package main

import (
	"fmt"
)

// Rectangle struct method without pointer
type Rectangle struct {
	Side float32
}

func (r Rectangle) GetArea() float32 {
	return r.Side * r.Side
}

func (r Rectangle) SetSide(side float32) {
	r.Side = side
}

// Circle struct method with pointer
type Circle struct {
	Radius float32
}

func (c *Circle) GetArea() float32 {
	return 3.14 * (c.Radius / 2) * (c.Radius / 2)
}

func (c *Circle) SetRadius(radius float32) {
	c.Radius = radius
}

// AddSide function to update side of rectangle, request rectangle model with pointer
func AddSide(r *Rectangle, add float32) {
	r.Side += add
}

// AddRadius function to update radius of circle, request circle model without pointer
func AddRadius(c Circle, add float32) {
	c.Radius += add
}

func main() {
	// initiate rectangle
	rect := Rectangle{Side: 2.5}

	// initiate circle
	circle := Circle{5}

	fmt.Println("area of rectangle before side edited:", rect.GetArea())
	rect.SetSide(4.5)
	fmt.Println("area of rectangle after side edited:", rect.GetArea())

	fmt.Println("area of circle before radius edited:", circle.GetArea())
	circle.SetRadius(10)
	fmt.Println("area of circle after radius edited", circle.GetArea())

	fmt.Println("side of rectangle before added", rect.Side)
	AddSide(&rect, 5)
	fmt.Println("side of rectangle after added", rect.Side)

	fmt.Println("radius of circle before added", circle.Radius)
	AddRadius(circle, 5)
	fmt.Println("radius of circle after added", circle.Radius)
}
