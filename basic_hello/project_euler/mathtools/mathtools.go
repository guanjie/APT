/* ericmath file for basic math calculations.
 */
package mathtools

import (
	NUM "../mathnumbers/"
	"fmt"
	"strconv"
)

func Sum(slice []int) int {
	result := 0
	for _, i := range slice {
		result += i
	}
	return result
}

/* @IsPalindrome: 放int的数，测试是不是palindrome.
 * USED By: 4
 */
func IsPalindrome(num int) bool {
	//准备
	str_num := strconv.Itoa(num)
	length := len(str_num)
	if length == 1 {
		return true
	}

	//加工
	for i := 0; i < length/2; i++ {
		if str_num[i] != str_num[length-1-i] {
			return false
		}
	}
	return true

	// 展示
}

/* @GetPrimeFactors: 放进去大数，算出来它的所有factors和多少次.
 * USED By: 3
 */
func GetPrimeFactors(big_num int) map[int]int {
	// 准备
	Primes := NUM.Primes
	map_primefactors := map[int]int{}

	// 加工
	for _, i := range Primes {
		// 如果能除尽这个prime, 继续除
		for big_num%i == 0 {
			// fmt.Println("Can divide: ", i)

			big_num = big_num / i
			// 如果没有数据就归零，如果有数据就加一
			map_primefactors[i] += 1
		}
		// 除到big_num < i. 不用继续除了.
		if big_num < i {
			// 展示
			// fmt.Println("The map_prime numbers are: ", map_primefactors)
			// fmt.Println("OUT, finished")
			return map_primefactors
		}
	}

	// TODO 到这边说明我们的Prime slice不够大.
	fmt.Println("WARNING: NOT COMPLETE, Need more prime numbers")
	fmt.Println("So far the map_prime numbers are: ", map_primefactors)
	return nil
}

/* @IsPrime: 查num是不是质数, 返回true/false
 * USED By: 3`(Getting Prime number slice)
 */
func IsPrime(num int) bool {
	// num <=0, 1, 2, 3
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

/* @GetEvenSlice: 查num是不是质数, 返回true/false
 * USED By: 2
 */
func GetEvenSlice(oldslice []int) []int {
	newslice := []int{}
	for _, i := range oldslice {
		if i%2 == 0 {
			newslice = append(newslice, i)
		}
	}
	return newslice
}
