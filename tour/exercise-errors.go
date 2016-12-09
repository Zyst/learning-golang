package exerciseerrors

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func compareEpsilon(epsilon, previous, current float64) bool {
	return !(math.Abs(previous-current) < epsilon)
}

func Sqrt(x float64) (float64, error) {

	// If the value is a negative number we return an error
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	aproximate := float64(1)
	previous := float64(0)

	// We iterate 10 times, or check if epislon condition clears
	for i := 0; i <= 10 && compareEpsilon(0.000000000000001, previous, aproximate); i++ {
		previous = aproximate
		aproximate = aproximate - ((math.Pow(aproximate, 2) - x) / (2 * aproximate))
	}

	return aproximate, nil
}

func main() {
	// Positive numbers should output as usual
	fmt.Println(Sqrt(2))

	// Negative numbers should throw an error
	fmt.Println(Sqrt(-2))

  fmt.Println(Sqrt(-999))
  fmt.Println(Sqrt(9))
}
