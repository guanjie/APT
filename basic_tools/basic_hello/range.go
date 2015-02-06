// Testing on range.go file for learning purpose.

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// Capable of finding the sum directly
	for _, num := range nums {
		sum += num
	}
	fmt.Println("The sum is", sum)

	// Capable of finding the right index directly
	for i, num := range nums {
		if num == 3 {
			fmt.Println("the index of 3 is:", i)
		}
	}

	// Key value map
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cao"}
	for k, v := range kvs {
		fmt.Println("{} -> {}\n", k, v)
	}

	vals := []string{"abc", "def"}
	for i, c := range vals {
		fmt.Println("i and c are: ", i, c)
	}

	// TRICK seperating the string to chars
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
