package main

import (
	// 	NUM "./mathnumbers"
	MATH "./mathtools"
	"fmt"
)

func main() {

	slice_primes := []int{}
	for i := 1; i < 200000; i++ {
		if MATH.IsPrime(i) {
			slice_primes = append(slice_primes, i)
		}
	}

	fmt.Println("length is", (slice_primes[10001]))
}
