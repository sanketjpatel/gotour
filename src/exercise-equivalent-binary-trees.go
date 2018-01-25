package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func RootWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go RootWalk(t1, ch1)
	go RootWalk(t2, ch2)
	
	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		
		if (ok1 != ok2) || (v1 != v2) {
			return false
		}
		
		if !ok1 {
			return true
		}
	}
}

func Leaf(x int) *tree.Tree {
	return &tree.Tree{nil, x, nil}
}

func main() {
	/*********
	 *   2   *
	 *  / \  *
	 * 1   3 *
         *********/
	t1 := &tree.Tree{Leaf(1), 2, Leaf(3)}
	
	/*********
         *    3  *
         *   /   *
         *  1    *
         *   \   *
         *    2  *
         *********/
	t2 := &tree.Tree{&tree.Tree{nil, 1, Leaf(2)}, 3, nil}
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
}

