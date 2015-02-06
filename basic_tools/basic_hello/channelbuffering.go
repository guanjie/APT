// channelbuffering.go file for learning purpose

package main

import (
    "fmt"
)

var p = fmt.Println

func main() {
    // TRICK channel buffer make sure the data is stored before executed
    c := make(chan string, 2)

    c <- "1st msg"
    c <- "2nd msg"
    p(<-c)
    p(<-c)

    c <- "3rd msg"
    p(<-c)

    var input string
    fmt.Scanln(&input)

}
