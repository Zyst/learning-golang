package indirection

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
  // This declaration requires an & when calling 'ScaleFunc'
	v := Vertex{3, 4}
  // This works the same on both
	v.Scale(2)
  // This function call requires a pointer
	ScaleFunc(&v, 10)

  // This declaration has an & passed directly, thus when calling ScaleFunc we
  // no longer need to pass the '&'
  p := &Vertex{3, 4}
  // Works the same in both
  p.Scale(2)
  // Doesn't need the '&'
  ScaleFunc(p, 10)

  // {60 80} &{60 80}
  fmt.Println(v, p)
}
