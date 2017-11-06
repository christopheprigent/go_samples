package main

import (
	"fmt"
)

func main() {

	r1 := Rectangle{leftX: 0, topY: 32, width: 42.2, height: 10.1}
	fmt.Printf("r : w: %g h: %g a: %g\n", r1.width, r1.height, r1.area())
}

// Rectangle struct for learning
type Rectangle struct {
	leftX  float32
	topY   float32
	height float32
	width  float32
}

func (r *Rectangle) area() float32 {
	return r.width * r.height
}
