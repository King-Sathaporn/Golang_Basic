package main

import (
	"fmt"
	"math"
)

//? Interface is a way to solve the polymorphism problem.
//? Polymorphism is the ability of an object to take on many forms.
//? Golang isn't a OOP language, so it doesn't have class.
//? But it has interface.
//? Interface will check error when compile time.

func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}

	fmt.Println("Area of rangetable: ", getArea(r))
	fmt.Println("Area of circle: ", getArea(c))
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}
