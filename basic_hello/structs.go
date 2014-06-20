// Structs.go file, for learning purpose.

package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
}

func main() {
    bob := person{"Bob", 39}

    p := fmt.Println
    p("Bob's age is:", bob.age)
    p(person{"Eric", 28})
    p(person{name: "humancool"})
    p(person{name: "humancool", age: 28})
    p(&person{name: "Sean", age: 50})

    s := person{"Sean", 50}
    sp := &s

    p(sp.name)
    p(s.age)
    sp.age = 100
    p(s.age)
}
