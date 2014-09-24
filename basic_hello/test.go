// testing.go for testing purpose
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func vals() (int, int) {
	return 3, 7
}

func main() {
	// test1 - return 2 values. We can at most pass in 2 values.
	fmt.Println(vals())

	//test2
	strs := []int{12, 34}
	for _, val := range strs {
		fmt.Println("val is:", val)
	}

	//test3 - strconv.itoa
	fmt.Println("This is " + strconv.Itoa(1+3))

	//test4 - appending strings, := here we did declaration and initialization
	s := make([]string, 0)
	s = append(s, "guanjie")
	s = append(s, "harvey")
	s = append(s, "niubi")
	fmt.Println(s)

	//test5 - Checking the number of chars in the string.
	fmt.Println("The string char length is: " + strconv.Itoa(len("361a585dc0b297027311b4d0123c55e4")))

	// test6 - test on the slices length.
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(len(letters))

	letters = append(letters, "e")
	fmt.Println(letters)
	fmt.Println(len(letters))
	fmt.Println(letters[len(letters)-1])

	// test 7, testing on adding 2 variables.
	v1 := 10
	v2 := 20
	v1, v2 = v1+v2, v1
	fmt.Println(v1, v2)

	// test 8, return value directly from the call
	fmt.Println(IntVector{100, 2, 3}.Sum())

	// test 9
	var million int = 1e6
	fmt.Println(million)
	fmt.Println(square(1.01))

	// test 10, rand.Perm playaround.
	fmt.Println("permuattion?", rand.Perm(10))
	candidates := rand.Perm(20)
	for val := range candidates {
		fmt.Println(val)
	}

	// test 11,
	ports := []int{1, 10, 2, 44, 20, 18}
	sort.Ints(ports)
	fmt.Println("the sort ports is: ", ports)

	// test 12
	fmt.Println("check it:", strings.Repeat("eric", 10))

	// test 13
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d is the value\n", i)
	}
}

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v { // blank identifier!
		s += x
	}
	return
}

func pump() chan int { // pump 1 to the channel ch
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func square(f float64) float64 {
	return f * f
}
