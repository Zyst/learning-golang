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
	// This sends values to the tree!
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	// We make two channels to walk trees
	ch1 := make(chan int)
	ch2 := make(chan int)

	// We send both trees to walk
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	firstCompare, secondCompare := WalkedTreeToArray(ch1), WalkedTreeToArray(ch2)

	// Can't do this with a single loop because items can be in different orders,
	for _, value := range firstCompare {
		for i := range secondCompare {
			if secondCompare[i] == value {
				// We "Delete" this value since it has matched once
				secondCompare = append(secondCompare[:i], secondCompare[i+1:]...)
				break
			}
		}
	}

	fmt.Println(firstCompare, secondCompare)

	// If we have 0 elements left they are the same
	return len(secondCompare) == 0
}

func WalkedTreeToArray(ch chan int) []int {
	result := make([]int, 10)
	for index := 0; index < 10; index++ {
		result[index] = <-ch
	}
	return result
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	fmt.Println(WalkedTreeToArray(ch))

	fmt.Println("\n\n")

	fmt.Println("Expect true: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Expect false: ", Same(tree.New(1), tree.New(2)))
}
