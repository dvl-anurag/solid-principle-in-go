package main

// Shape is an interface representing a geometric shape
type Shape interface {
	Area() float64
}

// Rectangle is a concrete implementation of the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is another concrete implementation of the Shape interface
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// CalculateTotalArea calculates the total area of multiple shapes
func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}
