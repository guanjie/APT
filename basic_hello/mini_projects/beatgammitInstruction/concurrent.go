package main

import (
	"fmt"
	"math/rand"
	// 	"os"
	// 	"strconv"
	"time"
)

// current number (shared between goroutines)
var num int = 0

// number of parts to split the computation into
const PARTS int = 10

func fillRandom(arr []int, c chan int, n int) {
	for i, _ := range arr {
		arr[i] = num
		num++

		// random sleep between 0-100 milliseconds
		dur := rand.Int63() % 100
		time.Sleep(time.Duration(dur) * time.Millisecond)
	}

	// write to the channel to signify that we're done
	c <- n
}

func goListen(c chan int, done chan bool) {
	rets := 0
	for rets < PARTS {
		fmt.Printf("Channel %d done\n", <-c)
		rets++
	}

	done <- true
}

func main() {
	// grab the seed from the command-line, if provided
	seed := 400

	rand.Seed(int64(seed))

	// make the array we're going to fill
	arr := make([]int, 888)

	// make a buffered channel
	c := make(chan int, PARTS)

	// the 'done' channel
	d := make(chan bool)
	go goListen(c, d)

	// slice up the array
	incr := len(arr) / PARTS
	for i := 0; i < PARTS; i++ {
		// create a slice for each segment
		slice := arr[i*incr : (i+1)*incr]

		// run each computation in it's own goroutine
		go fillRandom(slice, c, i)
	}

	// block until we're done
	<-d

	fmt.Println("All done!")
	fmt.Println(arr)
}
