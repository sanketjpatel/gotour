package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	seed1 := 0
	seed2 := 1
	
	return func() int {
		fibonacciNumber := seed1 + seed2
		
		seed1 = seed2
		seed2 = fibonacciNumber
		
		return fibonacciNumber
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", f())
	}
}

