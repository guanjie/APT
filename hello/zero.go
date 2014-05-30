package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 100
}

func main() {
    x := 5
    zero(&x)
    fmt.Println("x now is %v", x)
}
