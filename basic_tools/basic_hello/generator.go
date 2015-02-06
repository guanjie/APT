// idgenerator.go file for learning purpose.

package main

import (
	"fmt"
)

func idGenerator() chan int {
	ids := make(chan int)
	go func() {
		id := 0
		for {
			ids <- id
			id++
		}
	}()
	return ids
}

func main() {
	fmt.Println("Awesome Eric!")
	ids := idGenerator()

	id1 := <-ids
	id2 := <-ids
	id3 := <-ids
	id4 := <-ids

	fmt.Println(id1, id2, id3, id4)
}
