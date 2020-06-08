package main

import (
	"fmt"
	"math"
)

// Go does not have classes.
// A method is a function with a receiver argument.

type Vertex struct {
	X, Y float64
}

// Abs works on a copy of the receiver and cannot modify the original value.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
