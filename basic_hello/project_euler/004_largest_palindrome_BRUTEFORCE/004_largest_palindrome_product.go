/* 4_palindrome_largest_product
 *
 * A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit numbers.
 *
 * CONSTRAINTS: 三位数＊三位数的结果是正反读数
 * Result: 906609 = 993 * 913
 * HUMANEFFORT: Eyeballing, Bruteforce
 *
 */

package main

import (
	MATH "../mathtools"
	"fmt"
)

func main() {
	for i := 906609; i > 900000; i-- {
		if MATH.IsPalindrome(i) {
			fmt.Println("print: ", i)
			fmt.Println("FACTORS are: ", MATH.GetPrimeFactors(i))
			fmt.Println("\n")
		}
	}
}
