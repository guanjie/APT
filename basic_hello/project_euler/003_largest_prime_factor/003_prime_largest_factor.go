/* 3_prime_largest_factor
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 *
 * CONSTRAINTS: Prime factors
 * RESULT: 6857
 */

package main

import (
	// 	NUM "../mathnumbers/"
	MATH "../mathtools/"
	"fmt"
)

func main() {

	big_num := 600851475143
	fmt.Println("Prime factor map is: ", MATH.GetPrimeFactors(big_num))

}
