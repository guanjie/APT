// panic.go file, for learning purpose

package main

import (
    "fmt"
)

func main() {

    panic("There is a problem")

    _, err := fmt.Println("Print a line in here.")
    if err != nil {
        panic(err)
    }
}
