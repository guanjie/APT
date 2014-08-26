// testing.go for testing purpose
package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

func vals() (int, int) {
	return 3, 7
}

func main() {
	// test1 - return 2 values. We can at most pass in 2 values.
	fmt.Println(vals())

	//test2
	strs := []int{12, 34}
	for _, val := range strs {
		fmt.Println("val is:", val)
	}

	//test3 - strconv.itoa
	fmt.Println("This is " + strconv.Itoa(1+3))

	//test4 - appending strings, := here we did declaration and initialization
	s := make([]string, 0)
	s = append(s, "guanjie")
	s = append(s, "harvey")
	s = append(s, "niubi")
	fmt.Println(s)

	//test5 - Checking the number of chars in the string.
	fmt.Println("The string char length is: " + strconv.Itoa(len("361a585dc0b297027311b4d0123c55e4")))

	// test6 - test on the slices length.
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(len(letters))

	letters = append(letters, "e")
	fmt.Println(letters)
	fmt.Println(len(letters))
	fmt.Println(letters[len(letters)-1])

}
