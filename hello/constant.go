// constant.go file, for learnwing purpose

package main

import (
    "fmt"
)

var p = fmt.Println

func main() {
    const s = "testing string"
    p(s)

    const n = 100000
    p(int(n))
    // ATTENTION: const values can NOT be changed
    // n = 10
    // p(n)

    const boo_false = false
    p(boo_false)
    // boo_false = true
    // p(boo_false)

    const boo_true = true
    p(boo_true)

}
