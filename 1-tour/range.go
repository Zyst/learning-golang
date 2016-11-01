package range

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  // Two values are returned, first is index,
  // second is a copy of the value at that index
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }

  // Array with 0 '0's
  anotherPow := make([]int, 10)

  // We raise the pow just with i, skip the value
  for i := range anotherPow {
    anotherPow[i] = 1 << uint(i) // 2**i
  }

  // If we don't need the index we can use "_" instead
  for _, value := range anotherPow {
    fmt.Printf("%d\n", value)
  }
}
