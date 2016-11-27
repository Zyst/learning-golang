package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// I'll come back to this later, probably
func binaryConverter(decimal int) (binary []byte, err error) {
	if decimal < 0 {
		return nil, errors.New("binaryConverter: Can't convert negative numbers")
	}

	return result, nil
}

func main() {
	test, err := binaryConverter(16)

	if err != nil {
		log.Fatal(err)
	}

	// Lets convert 777 to binary!
	fmt.Println(test)
}
