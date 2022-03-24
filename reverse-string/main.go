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
	fmt.Println(reverseTwo("helicopter"))
	// last function on page for relevance of using kanji
	// these characters (and others in other languages/special characters) take up more than one byte
	fmt.Println(reverse("黒い犬"))
}

// good explanation on byte vs rune if unsure
// https://levelup.gitconnected.com/demystifying-bytes-runes-and-strings-in-go-1f94df215615
// in short bytes are uint8 and runes are int32 type

// reverse takes a string as an argument and returns the string reversed
func reverse(word string) string {
	// declare but not initialize variable
	var rStr string
	// this will iterate through the string starting from the final index position
	// go does not have [:-1] like python
	for i := len(word) - 1; i >= 0; i-- {
		// need to cast word to string because the characters in the string
		// are bytes (uint8) (unless using a range based loop like below - then they are runes (int32))
		rStr = rStr + string(word[i])
	}
	return rStr
}

// using strings package
// https://pkg.go.dev/strings#Builder
// in short - strings.Builder is more effecient than the above method because it minimizes memory copying and uses Write methods
// can use fmt.Fprintf with .Builder because it implements io.Writer interface
// don't copy a non-zero builder
// increased effeciency only noticeable on very large strings
func strBldRev(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}

// similar to above two methods - treating string as a list of characers and iterating through them
// adding each new character to the beginning of the string to achieve reverse order
func reverseTwo(word string) string {
	var rStr string
	// iterating through word in order instead of reverse order as above functions
	for i := 0; i < len(word); i++ {
		// add character and then rStr to achieve effect of adding each character to front of string
		// eg cat == c => ac => tac
		rStr = string(word[i]) + rStr
	}

	return rStr
}

// one thing to keep in mind
// strings in Go use bytes and some characters are more than one byte - for instance kanji characters and others
// best to use runes when dealing with special characters
// can always use runes whenever you want - this is probably the best option because it is more
// flexible than using bytes (supports international users better also)
func spCharRev(word string) string {
	var rStr string
	// using a range based loop to get runes
	// use blank identifier because index isn't necessary
	for _, r := range word {
		// same concept as function above
		// casting to string because rune and string aren't compatible
		rStr = string(r) + rStr
	}
	return rStr
}
