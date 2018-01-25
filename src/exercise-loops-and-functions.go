package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	err_margin := 1e-15
	z_prev := x
	z := x/2
	iter_count := 1
	
	for ; z-z_prev >= err_margin || z-z_prev <= -err_margin; iter_count++ {
		z_prev = z
		z -= (z*z - x) / (2*z)
	}
	fmt.Printf("\nIteration Count for finding square root = %v\n", iter_count)
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

