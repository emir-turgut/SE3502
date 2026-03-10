package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * 2
}

func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	rec := Rectangle{
		width:  12.4,
		height: 5.1,
	}

	circ := Circle{
		radius: 4,
	}

	printShapeInfo(rec)
	printShapeInfo(circ)
}
