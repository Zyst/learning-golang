package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Go tour exercise, this won't actually run because "golang.org/x/tour/wc"
// only runs online, to run it you must go to https://tour.golang.org/moretypes/23

func WordCount(s string) map[string]int {
	// We break the string fields into an array of strings
	words := strings.Fields(s)

	result := make(map[string]int)

	// We iterate through the words
	for _, v := range words {
		_, ok := result[v]

		// If a value exists we add 1 to its word count
		if ok {
			result[v]++
		} else {
			result[v] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}

/**
 * Sample output:

 PASS
 f("I am learning Go!") =
  map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"fox":1, "jumped":1, "over":1, "dog.":1, "quick":1, "brown":1, "the":1, "lazy":1, "The":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"Then":1, "another":1, "I":2, "ate":2, "a":1, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "man":1, "a":2, "plan":1, "canal":1, "panama.":1}
  
 */
