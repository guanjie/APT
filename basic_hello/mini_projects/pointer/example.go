// example.go file for learning purpose.
package main

import (
	"fmt"
)

func updateValue(valptr *float64) {
	*valptr = *valptr + 100
}

func main() {
	fmt.Println("Awesome Eric!")

	val := new(float64)
	updateValue(val)
	fmt.Println("val:", *val)
}
