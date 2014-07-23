// maps.go file for learning purpose

package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	// TRICK making map takes the make(map[string]int) function
	m := make(map[string]int)

	m["k1"] = 12
	m["k2"] = 123123

	//delete(m, "k1")
	p(m)

	_, prs := m["k2"]
	p(prs)

	n := map[string]int{}
	n["123123"] = 12312
	p(n)

	testing := []int{123123}
	testing = append(testing, 11)
	p(len(testing))

	var elements map[string]string
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
}
