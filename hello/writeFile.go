// writeFile.go file, for learning purpose

package main

import (
    "fmt"
    "os"
)

var p = fmt.Println

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    // TRICK after os.Open("pwd") can not write the file, only read it.
    f, err := os.Open("temp/testing.txt")
    check(err)
    p(f)
}
