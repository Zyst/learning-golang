package main

import (
	"fmt"
	"math"
)

func compareEpsilon(epsilon, previous, current float64) bool {
	return !(math.Abs(previous-current) < epsilon)
}

func Sqrt(x float64) float64 {
	aproximate := float64(1)
	previous := float64(0)

	// We iterate 10 times, or check if epislon condition clears
	for i := 0; i <= 10 && compareEpsilon(0.000000000000001, previous, aproximate); i++ {
		previous = aproximate
		aproximate = aproximate - ((math.Pow(aproximate, 2) - x) / (2 * aproximate))
	}

	return aproximate
}

func main() {
	fmt.Println("End:", Sqrt(4))
}
