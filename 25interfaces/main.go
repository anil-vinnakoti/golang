package main

import (
	"fmt"
	"math"
)

// Shape is an interface that requires implementing types to have Area and GetDimensions methods.
// Tip: Interfaces allow you to write flexible, reusable code.
type Shape interface {
	Area() float64
	GetDimensions() float32
}

// Rectangle is a struct that represents a rectangle with width and height.
type Rectangle struct {
	width, height float64
}

// Area calculates and returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// GetDimensions is a placeholder method required by the Shape interface.
// Tip: You can implement meaningful logic here, such as returning diagonal or perimeter if needed.
func (r Rectangle) GetDimensions() float32 {
	return 0
}

// Circle is a struct that represents a circle with a given radius.
type Circle struct {
	radius float64
}

// Area calculates and returns the area of the circle using the formula πr².
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Tip: Circle does not implement GetDimensions() method, so it does not fully implement Shape interface.
// Hence, you cannot pass Circle directly where Shape is expected unless you implement that method.

// calculateArea takes any Shape and returns its area.
// Tip: Using interfaces like this promotes polymorphism and clean abstractions.
func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	// Creating a Rectangle with width 6 and height 8
	rect := Rectangle{width: 6, height: 8}

	// Creating a Circle with radius 45
	circ := Circle{radius: 45}

	// Rectangle implements Shape interface, so this works
	fmt.Println("Rectangle area:", calculateArea(rect))

	// This line will cause a compile-time error because Circle does not implement GetDimensions().
	// Uncommenting the following line will raise an error until you implement GetDimensions() for Circle.

	// fmt.Println("Circle area:", calculateArea(circ))
}
