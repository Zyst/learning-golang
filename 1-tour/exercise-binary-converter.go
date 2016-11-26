package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func binaryConverter(decimal int) (binary []byte, err error) {
	if decimal < 0 {
		return nil, errors.New("binaryConverter: Can't convert negative numbers")
	}

	power := 0

	// First we find out how many "Powers of 2" we need
	// to fit in the passed decimal
	for {
		if math.Pow(2, float64(power)) > float64(decimal) {
			break
		}
		power++
	}

	result := make([]byte, power)

	// Restarting the meat of the for, I think I'll only grow more confused otherwise
	// When the soil is salted, move.

	return result, nil
}

func main() {
	test, err := binaryConverter(10)

	if err != nil {
		log.Fatal(err)
	}

	// Lets convert 777 to binary!
	fmt.Println(test)
}
