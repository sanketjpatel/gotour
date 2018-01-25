package main

import "fmt"

func addToChannel(ch chan int, x int) {
	ch <- x
}

func main() {
	ch := make(chan int, 1)
	go addToChannel(ch, 1)
	go addToChannel(ch, 2)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

