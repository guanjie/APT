package main

import (
	"fmt"
	//	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "buffered"
		ch <- "buffered1"
		ch <- "buffered2"
		// KEY is here to close the ch
		close(ch)
	}()

	// KEY-Secondary so that range or ch is possible, no blocking deadlock.
	for msg := range ch {
		fmt.Println("msg is: ", msg)
	}
}
