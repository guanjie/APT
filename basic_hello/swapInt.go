package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("Now we have x and y are: ", x, y)
}
