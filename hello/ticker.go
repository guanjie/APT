//ticker.go file for learning purpose
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 100)

    go func() {
        for t := range ticker.C {
            fmt.Println("The ticker ticks at: ", t)
        }
    }()

    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()

    fmt.Println("The ticker has now been stopped.")

}
