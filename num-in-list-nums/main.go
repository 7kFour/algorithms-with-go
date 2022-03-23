package main

import "fmt"

func main() {
	fmt.Println(numInList([]int{1, 2, 3, 4, 5}, 5))      // true
	fmt.Println(numInList([]int{3, 3, 3, 3, 3}, 5))      // false
	fmt.Println(numInList([]int{3, 5, 3, 5, 3}, 5))      // true
	fmt.Println(numInList([]int{4, 2, 22, -10, 8}, -10)) // true

	// testing for empty lists
	fmt.Println(numInList(nil, 5))     // false
	fmt.Println(numInList([]int{}, 5)) // false
}

// numInList will return true if num is in the list slice, and false otherwhise
func numInList(list []int, num int) bool {
	// think about what algo you want to use?
	// if sorted list could use search which is faster
	// lack of knowledge about the list means you must check every value
	// could use a map for this - or check every number (which is slower but more simple)

	// using blank identifier for index and iterating through the slice list
	for _, val := range list {
		// checking each value in list to see if it matches the number argument passed to the function
		if val == num {
			return true
		}
	}

	return false
}

// just because you see some test cases pass don't assume your code is correct
// make sure all of the applicable tests are passing
