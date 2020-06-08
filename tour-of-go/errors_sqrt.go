package main

import (
	"fmt"
	"math"
)

type ErrorNegativeSqrt float64

// Error for negative input for Sqrt. Must convert e to avoid infinite recursive loop.
func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrorNegativeSqrt(a)
	}
	xNext := 1.0
	xPrev := 0.0
	for math.Abs(xNext-xPrev) > 1e-8 {
		xPrev, xNext = xNext, xNext-(xNext*xNext-a)/(2*xNext)
		fmt.Println(xPrev, xNext)
	}
	return xNext, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
