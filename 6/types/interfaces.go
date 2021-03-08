package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func main() {
	r := Rect{width: 3, height: 4}
	measure(r)

	c := Circle{radius: 5}
	measure(c)
}
