package main

import "fmt"

func main() {
	defer fmt.Println(" World")
	defer fmt.Println("Eric here")

	fmt.Println("Hello")
}
