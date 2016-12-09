package exerciseloopsandfunc

import "testing"
import "math"

func TestSqrt(t *testing.T) {

	type args struct {
		input float64
		want float64
	}

	tests := []args {
		{ 4, 2 },
		{ 9, 3 },
		{ 7, math.Sqrt(7) },
		{ 192, math.Sqrt(192) },
	}

	for _, item := range tests {

		// If our result is different from our desired result
		if got := Sqrt(item.input); got != item.want {
			t.Errorf("Error: Wanted %v, got %v", item.want, got)
		}
	}
}
