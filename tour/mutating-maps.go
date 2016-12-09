package mutatingmaps

import (
	"fmt"
)

func main() {
  m := make(map[string]int)

  // We can insert an element by declaring it
  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  // Values are mutable
  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  // We can delete by doing the following
  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"]) // 0

  // To test if a key is present we do a two value assignment
  v, ok := m["Answer"]

  // If a key were in m "ok" would be true, since there isn't it's false
  fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}
