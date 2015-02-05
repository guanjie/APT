/* multiplies of 3 and 5
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, 
 * we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 * CONSTRAINTS: BELOW 1000
 * RULE: get most generic solution possible
 * RESULT: 233168
 */

package main

import (
	"fmt"
)

func main() {
	// 准备: 合并两个map
	map1 := getMultiples(999, 3) // BELOW 1000
	map2 := getMultiples(999, 5) // BELOW 1000
	combinedmap := combineMaps(map1, map2)

	// 加工: 加和map里面数据
	sum := 0
	for k, _ := range combinedmap {
		sum += k
	}

	// 展示: 打印最后map和加和
	fmt.Println(combinedmap)
	fmt.Println("Total sum of the mulfiples of 3 and 5 below 1000 is", sum)
}

// denominator is 分母, mapping只回true的
func getMultiples(num, denom int) map[int]bool {
	mapping := make(map[int]bool)

	// 分母是0时
	if denom == 0 {
		// 不会任何值
		return
	}

	// 分母不是0时，遍历找到所有相关数字
	for i := 1; i <= num; i++ {
		if i%denom == 0 {
			fmt.Printf("%d is divisible by %d\n", i, denom)
			mapping[i] = true
		}
	}
	return mapping
}

// 合并 map[int]bool
func combineMaps(map1, map2 map[int]bool) map[int]bool {
	for k, v := range map1 {
		map2[k] = v
	}
	return map2
}
