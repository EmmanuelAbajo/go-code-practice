package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Rectangle struct
type Rectangle struct {
	length float64
	breadth float64
}

// Area of Circle implementing shape interface
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area of Circle implementing shape interface
func (r *Rectangle) Area() float64 {
	return r.length * r.breadth
}

func main() {
	var shapes []Shape
	shapes = []Shape{ &Circle{10}, &Rectangle{4,5} }
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	fmt.Printf("Total Area: %v\n", total)
}