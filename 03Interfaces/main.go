package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func measure(s Shape) {
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func main() {
	r := Rectangle{3, 4}
	c := Circle{5}

	measure(r)
	measure(c)

}
