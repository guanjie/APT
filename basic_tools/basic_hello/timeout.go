// timeout.go file for learning purpose.

package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(time.Duration(2) * time.Second)

	for i := 0; i < 10; i++ {
		// Sleep 10 cycles
		time.Sleep(time.Duration(i) * time.Second)

		select {
		case <-timeout:
			fmt.Println("Too late man, after 2 seconds")
		}
	}
}
