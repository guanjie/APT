// readfiles.go file for learning purpose

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

var p = fmt.Println

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    data, err := ioutil.ReadFile("temp/testing.txt")
    check(err)
    p(string(data))
    p("=================\n")

    f, err := os.Open("temp/testing.txt")
    check(err)
    // File reading could be further checked
    p(f)

}
