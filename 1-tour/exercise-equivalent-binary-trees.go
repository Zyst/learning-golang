package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Quasi recursive way to walk the tree
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	// Quasi recursive way to walk the tree
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	// This is gonna run regarldess
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// This actually needs to use Walk, so no cigar
	return t1.Value == t2.Value
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for index := 0; index < 10; index++ {
		fmt.Println(<-ch)
	}
}
