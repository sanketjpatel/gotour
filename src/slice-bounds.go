package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	
	s[0] = 6
	
	fmt.Println(primes)
}

