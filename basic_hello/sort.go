// sort.go file for learning purpose.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Awesome Eric!")
	arr := []string{"c", "ddd", "aa"}
	sort.Strings(arr)
	fmt.Println(arr)

	new_string := strings.Join(arr, "")
	fmt.Println(new_string)
}
