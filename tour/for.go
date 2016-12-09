package forpack

import (
	"fmt"
)

func main() {
	sum := 0

	// Standard for notation
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	// This is the equivalent of a "while"
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
