package methodspointersexplained

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}

  // If we don't have the & here we get:
  // cannot use v (type Vertex) as type *Vertex in argument to Scale
  Scale(&v, 10)
  fmt.Println(Abs(v))
}
