package bufferedchannels

import (
	"fmt"
)

func main() {
  // A buffered channel is initialized with the buffer length
  // as the second argument, 2 in this case.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
  // For instance if we add this:
  // ch <- 3
  // We would get:
  // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
