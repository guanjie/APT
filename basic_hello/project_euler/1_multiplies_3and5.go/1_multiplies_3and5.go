// multiplies of 3 and 5.go file for learning purpose.
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
// ATTN: BELOW 1000
// Rule - get most generic solution possible

// Result: 233168
package main

import (
	"fmt"
)

func main() {
	map1 := getMultiples(999, 3) // BELOW 1000
	map2 := getMultiples(999, 5) // BELOW 1000
	combinedmap := combineMaps(map1, map2)
	sum := 0
	for k, _ := range combinedmap { // Adding all keys in the combined map
		sum += k
	}

	fmt.Println(combinedmap)
	fmt.Println("Total sum of the mulfiples of 3 and 5 below 1000 is", sum)
}

func getMultiples(num, denom int) map[int]bool {
	mapping := make(map[int]bool)

	if denom == 0 { // If the denominator is 0
		return mapping
	}

	for i := 1; i <= num; i++ {
		if i%denom == 0 {
			fmt.Printf("%d is divisible by %d\n", i, denom)
			mapping[i] = true
		}
	}
	return mapping
}

func combineMaps(map1, map2 map[int]bool) map[int]bool {
	for k, v := range map1 {
		map2[k] = v
	}
	return map2
}
