// Given a number N, prinout out all the numbers from 1 to N
// when a number is divisble by 3 print "Fizz" instead of the number
// when a number is divisible by 5 print "Buzz" instead of the number
// when a number is divisible by both 3 and 5 print "Fizz Buzz" instead of the number
// Ex. 1, 2, Fizz, 4, Buzz etc... separated by comma and space
// no comma after final number of range
package main

import "fmt"

func main() {
	fizzBuzz(100)
}

// putting everything in one function for readability
// could also split this out in to 2 or more functions
func fizzBuzz(n int) {
	// start at 1 in order to avoid dividing in to 0
	// and getting "Fizz Buzz" before 1
	// check for %15 first otherwise first condition reached will be fulfilled
	// eg if checking %3 %5 %15 each n divisble by 15 will print out "Fizz"
	// because they will also be divisible by 3 and 3 was reached first

	// using a switch statement - i find them very readable
	// a few different examples of how switch statements work https://gobyexample.com/switch
	// can use %v for a default format specifier if you want

	// reasonably simple way to handle adding a newline and a no comma or space after the final n
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("%s", "Fizz Buzz")
		case i%5 == 0:
			fmt.Printf("%s", "Buzz")
		case i%3 == 0:
			fmt.Printf("%s", "Fizz")
		default:
			fmt.Printf("%d", i)
		}

		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Print("\n")
}

// example of if/else version - this works as well
// use w/e you prefer or your organization dictates

// for i := 1; i <= n; i++ {
// 	if i%15 == 0 {
// 		fmt.Printf("%v, ", "Fizz Buzz")
// 	} else if i%5 == 0 {
// 		fmt.Printf("%v, ", "Buzz")
// 	} else if i%3 == 0 {
// 		fmt.Printf("%v, ", "Fizz")
// 	} else {
// 		fmt.Printf("%v, ", i)
// 	}
// }
