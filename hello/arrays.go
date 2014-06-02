//Arrays.go file for learning purpose. Check.
package main

import (
    "fmt"
)

func main() {
    // After doing var on the array you can modify the inner values from the arr.
    var arr [5]int
    fmt.Println("empty arr is:", arr)

    arr[3] = 34
    fmt.Println("after changing the 4th value to 34 the arr became: ", arr)

    fmt.Println("length of the array is: ", len(arr))

    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("New arr2 is: ", arr2)

    // Two dimention arrays, you will need to define the array itself.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("New twoD matrix looks like this: ", twoD)
}
