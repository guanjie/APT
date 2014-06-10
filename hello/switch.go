// switch.go file, for learning purpose

package main

import (
    "fmt"
    "time"
)

var p = fmt.Println

func main() {

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        p("it's the weekend today!")
    default:
        p("it's the weekday today.")
    }

    // TRICK switch arguments only choose one condition out of the many
    switch {
    case time.Now().Hour() > 12:
        p("Good afternoon everybody")
    case time.Now().Minute() < 30:
        p("We are still in the first half of this hour.")
    }

}
