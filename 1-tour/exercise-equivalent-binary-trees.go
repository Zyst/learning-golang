package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

  // Do this somehow?
  ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  return t1.Value == t2.Value
}

func main() {
  ch := make(chan int)

  go Walk(tree.New(1), ch)
}

