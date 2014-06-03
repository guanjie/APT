// recursion.go file, for learning purpose

package main

import (
    "fmt"
)

func factorial(n int) int {
    if n <= 0 {
        return 1
    } else {
        return n * factorial(n-1)
    }
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
        fmt.Println(factorial(i))
    }
}
