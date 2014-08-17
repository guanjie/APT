// slices.go file for learning purposes

package main

import (
	"fmt"
)

var p = fmt.Println

func main() {

	// TRICK, s := make([]string, 3), need to specify the length of the slices
	s := make([]string, 6)
	// 	s := []string{}

	s[0] = "First string"
	s[1] = "Second string"
	s = append(s, "fourth")
	s = append(s, "5", "6")

	c := make([]string, len(s))
	copy(c, s)
	s[0] = "changed"
	p(s)
	p(c)

	for index, val := range s {
		fmt.Println("value index is: ", index)
		fmt.Println("value val is: ", val)
	}
}
