package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Circle struct
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define a Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Main function
func main() {
	var s Shape

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	// Assign Circle to the interface
	s = c
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())

	// Assign Rectangle to the interface
	s = r
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
