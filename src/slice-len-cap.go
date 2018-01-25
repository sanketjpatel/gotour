package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	s := primes[:]
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[2:6]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	
	printSlice(primes)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

