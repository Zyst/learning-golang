package nilslices

import (
	"fmt"
)

func main() {
  // The zero value of a slice is nil
  var s []int
  fmt.Println(s, len(s), cap(s))

  if s == nil {
    fmt.Println("nil!")
  }
}
