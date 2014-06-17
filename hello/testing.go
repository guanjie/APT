// testing.go for testing purpose
package main

import (
    "fmt"
    "os"
)

var p = fmt.Println

func vals() (int, int) {
    return 3, 7
}

func main() {
    fmt.Println(vals())
    strs := []int{12, 34}
    for _, val := range strs {
        fmt.Println("val is:", val)
    }
}
