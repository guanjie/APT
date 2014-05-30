// pointer.go file for learning purpose

package main

import (
    "fmt"
)

func zeroval(ival int) {

    ival = 0
}

func zeroptr(iptr *int) {

    *iptr = 0
}

func main() {
    p := fmt.Println
    i := 100
    p("Initial value is: ", i)

    zeroval(i)
    p("after zeroval func i is: ", i)

    zeroptr(&i)
    p("After zeroptr func i is: ", i)

    p("point val of i is: ", &i)
}
