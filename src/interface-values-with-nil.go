package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	//
	fmt.Println(t.S)
}

func main() {
	var i I

	var t T
	i = t
	describe(i)
	i.M()

	i = T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

