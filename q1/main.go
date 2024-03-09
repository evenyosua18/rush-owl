package main

import (
	"fmt"
)

type Rectangle struct {
	Side float32
}

func (r Rectangle) GetArea() float32 {
	return r.Side * r.Side
}

func (r Rectangle) SetSide(side float32) {
	r.Side = side
}

type Circle struct {
	Radius float32
}

func (c *Circle) GetArea() float32 {
	return 3.14 * (c.Radius / 2) * (c.Radius / 2)
}

func (c *Circle) SetRadius(radius float32) {
	c.Radius = radius
}

func AddSide(r *Rectangle) {
	r.Side += 5
}

func AddRadius(c Circle) {
	c.Radius += 5
}

func main() {
	rect := Rectangle{Side: 2.5}

	fmt.Println(rect.GetArea())
	rect.SetSide(4.5)
	fmt.Println(rect.GetArea())

	circle := Circle{5}

	fmt.Println(circle.GetArea())
	circle.SetRadius(10)
	fmt.Println(circle.GetArea())

	fmt.Println(rect.Side)
	AddSide(&rect)
	fmt.Println(rect.Side)

	fmt.Println(circle.Radius)
	AddRadius(circle)
	fmt.Println(circle.Radius)
}
