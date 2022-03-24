// reverse will return the provided word in reverse order
// eg reverse("truck") => "kcurt"
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("tractor"))
	fmt.Println(strBldRev("StringBuilder"))
}

// reverse takes a string as an argument and returns the string reversed
func reverse(word string) string {
	// declare but not initialize variable
	var rStr string
	// this will iterate through the string starting from the final index position
	// go does not have [:-1] like python
	for i := len(word) - 1; i >= 0; i-- {
		// using runes like this will give a a unicode code point (an int32)
		// so its necessary to cast to string when rebuilding rStr
		rStr = rStr + string(word[i])
	}
	return rStr
}

// using strings package
// https://pkg.go.dev/strings#Builder
// in short - strings.Builder is more effecient than the above method because it minimizes memory copying and uses Write methods
// can use fmt.Fprintf with .Builder because it implements io.Writer interface
// don't copy a non-zero builder
func strBldRev(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}

