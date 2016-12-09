package channels

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// Slice of First half of s
	go sum(s[:len(s)/2], c)
	// Slice of Latter half of s
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // Receive from c

	fmt.Println(x, y, x+y)
}
