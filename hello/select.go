package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "One"
            time.Sleep(time.Second * 1)
        }
    }()

    go func() {
        for {
            c2 <- "Two"
            time.Sleep(time.Second * 2)
        }
    }()

    for i := 0; i < 100; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Msg1 is: ", msg1)
        case msg2 := <-c2:
            fmt.Println("Msg2 is: ", msg2)
        }
    }
}
