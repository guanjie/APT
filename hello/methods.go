// methods.go file for learning purpose

package main

import (
    "fmt"
)

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

// TRICK if pass in addressed struct, then we can change the builtin val
func (r *rect) perim() int {
    // ATTN here
    r.width = 1
    r.height = 2
    return 2 * (r.width + r.height)
}

func main() {

    rect1 := rect{width: 10, height: 20}
    p := fmt.Println

    p(rect1)
    p("area is ", rect1.area())
    p("perimeter is ", rect1.perim())
    p("area is ", rect1.area())
    p(rect1.width)
    p(rect1.height)

}
