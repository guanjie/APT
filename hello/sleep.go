package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 10; i++ {
        //     select {
        //        case <-time.After(time.Second):
        //           fmt.Println("this is after one second number: ", i)
        //      }

        time.Sleep(time.Second)
        fmt.Println("This is after one second and the number is", i)
    }
}
