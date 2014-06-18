// Changing the value of a val directly through its address

package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 100
}

func main() {
    x := 5
    zero(&x)
    fmt.("x now is %v", x)
}
