// given a list of numbers, return the sum of the list

package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))    // 15
	fmt.Println(sum([]int{3, 3}))             // 6
	fmt.Println(sum([]int{3, 5, 3, 5, 3}))    // 19
	fmt.Println(sum([]int{4, 2, 22, -10, 8})) // 26

	// let's just return 0 for empty lists
	fmt.Println(sum(nil))     // 0
	fmt.Println(sum([]int{})) // 0
}

// sum will sum up all the numbers passed in and return the result
func sum(numbers []int) int {
	// variable to add and store value
	s := 0

	// iterate through passed list and add value to s variable
	// to return a total
	for _, val := range numbers {
		s += val
	}

	return s
}
