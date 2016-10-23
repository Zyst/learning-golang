package pointers

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         // Point to i
	fmt.Println(*p) // Read i though the pointer
	*p = 21         // Set i through the pointer
	fmt.Println(i)  // See the new value of i

	p = &j       // Point to j
	*p = *p / 37 // Divide j through the pointer
	fmt.Println(j)
}
