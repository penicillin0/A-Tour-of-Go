package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	b := 0.0
	for z-b >= 0.01 {
		b = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(b, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
