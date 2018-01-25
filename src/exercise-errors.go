package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

func Sqrt(x float64) (float64, error) {
	err_margin := 1e-15
	z_prev := x
	z := x/2
	
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	for z-z_prev >= err_margin || z-z_prev <= -err_margin {
		z_prev = z
		z -= (z*z - x) / (2*z)
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

