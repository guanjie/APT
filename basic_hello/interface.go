// interface.go file, for learning purpose, new

package main

import (
    "fmt"
    "math"
)

var p = fmt.Println

type geometry interface {
    area() float64
    perim() float64
}

type square struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (s square) area() float64 {
    return s.width * s.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// TRICK, as if type square struct is g geometry
func measure(g geometry) {
    p("area is: ", g.area())
    p("perim is: ", g.perim())
}

func main() {
    s := square{width: 3, height: 4}
    c := circle{radius: 5}

    measure(s)
    measure(c)
}
