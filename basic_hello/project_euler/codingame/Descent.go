package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	for {
		var SX, SY int
		fmt.Scan(&SX, &SY)
		var i_highest int
		var MH_highest int = 0

		for i := 0; i < 8; i++ {
			// MH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
			var MH int
			fmt.Scan(&MH)

			if MH > MH_highest {
				MH_highest = MH
				i_highest = i
			}
		}
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		if SX == i_highest {
			fmt.Println("FIRE")
		} else {
			fmt.Println("HOLD")
		}
	}
}
