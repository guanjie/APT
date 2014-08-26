// else.go file for learning purpose.
package demo

import (
//	"fmt"
)

func comparedToZero(num int) int {
	if num > 0 {
		return 1
	} else if num == 0 {
		return 0
	} else {
		return -1
	}
}
