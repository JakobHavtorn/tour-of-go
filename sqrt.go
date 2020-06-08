package main

import (
	"fmt"
	"math"
)

// Sqrt computes the square root of a number.
// We want to find the square root of a number a, x=sqrt(a), i.e. solve x^2=a.
// Newtons method finds the root of f(x)=0 by approximating f(x) as its tangent
// line f(x_n)+f'(x_n)(x−x_n), leading to an improved guess x_n+1 from the root
// of the tangent x_n+1 = x_n - f(x_n) / f'(x_n)
// Letting f(x) = x^2 - a, we can solve the square root problem x^2=a.
// We have f'(x)=2x so the update equation is x = x - (x^2 - a) / 2x
func Sqrt(a float64) float64 {
	x := 1.0
	xTmp := x
	eps := 1.0
	for eps > 1e-15 {
		xTmp -= (x*x - a) / (2 * x)
		eps = math.Abs(x - xTmp)
		fmt.Println(x, xTmp, eps)
		x = xTmp
	}
	return x
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
