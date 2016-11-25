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

	// We go decreasingly, from the last element to the first element
	for index := len(result); index > 0; index-- {
    // If the current Power is larger than the decimal
		if i := math.Pow(2, float64(index)); i < float64(decimal) {
      // We subtract the current power from decimal
      fmt.Printf("Power val is: %v, Decimal is: %v\n", i, decimal)
      decimal -= int(i)
      fmt.Printf("Power val is: %v, Decimal is: %v\n", i, decimal)
      fmt.Printf("Index is: %d\n\n", index-1)

      // This bit should be a
			result[index-1] = 1
		} else {
      fmt.Printf("Power val is: %v, Decimal is: %v\n", i, decimal)
      fmt.Printf("Index is: %d\n\n", index-1)
      result[index-1] = 0
    }
	}

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
