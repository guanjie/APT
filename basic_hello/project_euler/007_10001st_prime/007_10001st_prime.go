/* 7_10001st_prime
 *
 * By listing the first six prime numbers: 2,3,5,7,11 and 13, we can see that the 6th is 13
 * What is the 10001st prime number?
 *
 * CONSTRAINTS: Prime Number slice
 * RESULT: 104743
 */

package main

import (
	// 	NUM "./mathnumbers"
	MATH "../mathtools"
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
