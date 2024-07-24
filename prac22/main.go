package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, length float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.length
}

type Shape interface {
	Area() float64
}

type AnyShape struct {
	width float64
}

func (a AnyShape) Area() float64 {
	return a.width * a.width
}

func main() {
	var i Shape
	i = &Rectangle{
		width:  7.1,
		length: 8.1,
	}
	fmt.Println(i.Area())

	i = AnyShape{
		width: 9.1,
	}
	fmt.Println(i.Area())

}
