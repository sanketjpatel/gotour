package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Hour()
	switch {
	case x < 12:
		fmt.Println("Good morning!")
	case x < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

