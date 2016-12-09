package functions

import "testing"

type addpair struct {
	x int
	y int
	expected int
}

var tests = []addpair {
	{ 2, 4, 6 },
	{ 1248, 1752, 3000 },
	{ 5, 8, 13 },
}

func TestAdd(t *testing.T) {

	// We check a value mnually
	six := add(2, 4)

	if six != 6 {
		t.Errorf("Sum incorrect, expected 6, got %d", six)
	}

	// We check a value mnually
	threeThousand := add(1248, 1752)

	if threeThousand != 3000 {
		t.Errorf("Sum incorrect, expected 3000, got %d", threeThousand)
	}

	// This is how we would create a test 'suite' instead to check many different values
	for _, pair := range tests {
		result := add(pair.x, pair.y)

		if result != pair.expected {
			t.Errorf("Sum incorrect, expected %d, got %d", pair.expected, result)
		}
	}
}
