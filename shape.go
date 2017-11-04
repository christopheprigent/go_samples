package main

import (
	"fmt"
	"math"
)

func main() {

	r1 := BetterRectangle{width: 42.2, height: 10.1}
	c1 := BetterCircle{radius: 11}

	fmt.Printf("rect area : %g\n", getArea(r1))
	fmt.Printf("circle area : %g\n", getArea(c1))
}

// Shape struct for learning
type Shape interface {
	area() float64
}

// BetterRectangle struct for learning
type BetterRectangle struct {
	height float64
	width  float64
}

// BetterCircle struct for learning
type BetterCircle struct {
	radius float64
}

func (r BetterRectangle) area() float64 {
	return r.width * r.height
}

func (c BetterCircle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// getArea for all shapes
func getArea(s Shape) float64 {
	return s.area()
}
