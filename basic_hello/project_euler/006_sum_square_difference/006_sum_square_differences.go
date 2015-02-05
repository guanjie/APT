/* 6_sum_square_difference
 *
 * The sum fo the squares of the first ten natural number is,
 * 1^2 + 2^2 + 3^2 + 4^2 + 5^2 + 6^2... + 10^2 = 385
 * The square of the sum of the first ten numbers is:
 * (1+2+3...+10)^2 = 55^2 = 3025
 * Hence the difference is 3025-385=2640
 * Find the difference between the sum of the squares of the first 100 natural numbers
 * and their square of the sum.
 *
 * CONSTRAINTS: Integers
 * Result: 25164150
 *
 */

package main

import (
	// 	MATH "../mathtools"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Abs diff between SumOfSquare and SquareOfSum is: ", math.Abs(SquareOfSum(100)-SumOfSquare(100)))
}

func SumOfSquare(num float64) float64 {
	sum := 0.0
	if num < 0 {
		fmt.Println("WARNING, CANT BE < 0")
		return 0
	}

	// 加工
	for i := 0.0; i <= num; i++ {
		sum += i * i
	}
	return sum
}

func SquareOfSum(num float64) float64 {
	sum := 0.0
	if num < 0 {
		fmt.Println("WARNING, CANT BE < 0")
		return 0
	}

	// 加工
	for i := 0.0; i <= num; i++ {
		sum += i
	}
	sum = sum * sum
	return sum
}
