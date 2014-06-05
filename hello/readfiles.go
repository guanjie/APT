// readfiles.go file for learning purpose

package main

import (
    "fmt"
)

var p = fmt.Println

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    p("Testing Testing.")
}
