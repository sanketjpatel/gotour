package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 0
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

