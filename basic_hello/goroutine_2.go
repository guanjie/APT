// goroutine2.go file for testing purpose.
package main

import (
	"fmt"
	"time"
)

func computation(n int) {
	time.Sleep(time.Second * time.Duration(n))
	fmt.Printf("%d millisenconds elapsed.\n", n)
}

func main() {

	go computation(1)
	go computation(3)
	go computation(5)

	// WHY we need another non go function to get the script rolling.
	computation(5)

}
