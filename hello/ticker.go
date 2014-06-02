// ticker.go file for learning purpose
// Didnt get it all, I don't see where we started the ticker
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 100)

    go func() {
        time.Sleep(time.Millisecond * 400)
        for t := range ticker.C {
            fmt.Println("The ticker ticks at: ", t)
        }
    }()

    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()

    fmt.Println("The ticker has now been stopped.")

}
