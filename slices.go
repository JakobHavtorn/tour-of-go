package main

import (
	"fmt"
	"math"

	"code.google.com/p/go-tour/pic"
)

// Pic returns a matrix of function evaluations
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx)
	fmt.Println(len(result), cap(result))
	for x := 0; x < dx; x++ {
		result[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			result[x][y] = uint8(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
