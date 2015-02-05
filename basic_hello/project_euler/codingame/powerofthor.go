package main

import "fmt"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// LX: the X position of the light of power
	// LY: the Y position of the light of power
	// TX: Thor's starting X position
	// TY: Thor's starting Y position
	var LX, LY, TX, TY int
	fmt.Scan(&LX, &LY, &TX, &TY)

	for {
		// E: The level of Thor's remaining energy, representing the number of moves he can still make.
		var E int
		fmt.Scan(&E)

		// First walk in mid-angles
		if TX < LX && TY < LY {
			fmt.Println("SE")
			TX += 1
			TY += 1
		} else if TX < LX && TY > LY {
			fmt.Println("NE")
			TX += 1
			TY -= 1
		} else if TY < LY && TX > LX {
			fmt.Println("SW")
			TX -= 1
			TY += 1
		} else if TX > LX && TY > LY {
			fmt.Println("NW")
			TX -= 1
			TY -= 1
		} else
		// Then walk in direct lines
		if TY > LY {
			fmt.Println("N")
		} else if TX > LX {
			fmt.Println("W")
		} else if TY < LY {
			fmt.Println("S")
		} else if TX < LX {
			fmt.Println("E")
		}
	}
}
