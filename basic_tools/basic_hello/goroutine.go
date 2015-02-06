// goroutine.go file for learning purpose

package main

import (
    "fmt"
    "time"
)

var p = fmt.Println

func f(from string) {
    for i := 0; i < 8; i++ {
        p(from, ":", i)
        time.Sleep(time.Millisecond * 100)
    }
}

func main() {

    // TRICK, inside for loop operates every goroutine.
    go f("direct")

    go f("synchronous")

    go f("come on, another one")

    go f("the fourth one, make description longer")

    var input string
    fmt.Scanln(&input)

}
