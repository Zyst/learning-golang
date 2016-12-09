package slicelencap

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)

	// Try to extend over its capacity
	// s = s[:10] // panic: runtime error: slice bounds out of range
}

func printSlice(s []int) {
	fmt.Printf("len(gth)=%d cap(acity)=%d %v\n", len(s), cap(s), s)
}
