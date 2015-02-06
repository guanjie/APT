// epoch.go file for learning purpose.

package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()
    p := fmt.Println

    p(now)
    p(now.Unix())
    p(now.UnixNano())

    p(time.Unix())

}
