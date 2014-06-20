// channel.go file, for learning purpose

package main

import (
    "fmt"
    "time"
)

var p = fmt.Println

func ping(c chan string) {

    for i := 0; ; i++ {
        c <- "ping"
        p("this is the nth ping:", i)
    }
}

func main() {

    c := make(chan string)
    // ping is a for loop open all the time
    go ping(c)

    go func() {
        for {
            time.Sleep(time.Millisecond * 100)
            msg := <-c
            p(msg)
        }
    }()

    var input string
    fmt.Scanln(&input)
}
