// 3_prime_largest_factor.go for Euler learning purpose.
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

// Result: 6857

package main

import (
	"fmt"
	// 	"math"
)

func main() {

	// 	for i := int(math.Sqrt(600851475143)); i > 1; i-- { // Check the squareroot of the number
	for i := 600851475143 / 2; i > 3; i-- { // This should be the complete way. This is done and with charm within 2 hrs. 6857 is the correct answer.
		if 600851475143%i == 0 {
			if IsPrime(i) {
				fmt.Println(i)
				break
			}
		}
	}
}

// only need to divide the prime numbers to check if the number is prime.
func IsPrime(num int) bool {
	// Corner cases when num is <=0, 1, 2, 3
	if num <= 0 { // if num is less than or equal to 0
		return false
	}
	if num == 1 || num == 2 || num == 3 { // if num is 1,2,3
		return true
	}

	// starting from 4 we do the checking, since 4 has at least 1 odd number 3
	for i := 2; i <= num/2; i++ { // ATTN: only need to go through the prime numbers
		if num%i == 0 {
			return false
		}
	}
	return true // if all not divisible, then return isPrime.
}
