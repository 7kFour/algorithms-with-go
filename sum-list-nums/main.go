// given a list of numbers, return the sum of the list

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))    // 15
	fmt.Println(sum([]int{3, 3}))             // 6
	fmt.Println(sum([]int{3, 5, 3, 5, 3}))    // 19
	fmt.Println(sum([]int{4, 2, 22, -10, 8})) // 26

	// let's just return 0 for empty lists
	fmt.Println(sum(nil))     // 0
	fmt.Println(sum([]int{})) // 0

	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("Recursive function starts now")
	fmt.Println(strings.Repeat("#", 10))

	// recursive version
	fmt.Println(rSum([]int{1, 2, 3, 4, 5}))    // 15
	fmt.Println(rSum([]int{3, 3}))             // 6
	fmt.Println(rSum([]int{3, 5, 3, 5, 3}))    // 19
	fmt.Println(rSum([]int{4, 2, 22, -10, 8})) // 26

	// let's just return 0 for empty lists
	fmt.Println(rSum(nil))     // 0
	fmt.Println(rSum([]int{})) // 0
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

// recursive version of above
// this isn't a good idea for large lists
// this is a less clear solution to the one above but good to help illustrate recursion
// best to understand when recursion is the best option
func rSum(numbers []int) int {

	// make sure there is a value in the first spot
	// ie make sure the list isn't empty
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + rSum(numbers[1:])
}

// beware stack overflow
// ex
// rSum(3, 4, 5)
// 	3 + rSum(4, 5)
// 		4 + rSum(5)
// 			5 + rSum()
// 				0 <== this is where our if statement will cause the function to stop recursion
