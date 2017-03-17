package main

import (
	"fmt"
	"math"
)

type Frame struct {
	X float32
	Y float32
}

type Circle struct {
	Center Frame
	Radius float32
}

func (c Circle) String() string {
	return fmt.Sprintf("Center(%.2f, %.2f), Radius %.2f", c.Center.X, c.Center.Y, c.Radius)
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	c1 := Circle{
		Center: Frame{X:-10, Y:20},
		Radius: 20,
	}

	fmt.Println(c1)
	fmt.Println("Area is", c1.Area())
}