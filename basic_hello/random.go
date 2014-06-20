// This is the random go file.
package main

import "fmt"
import "math/rand"

func main() {
    for i := 0; i < 50; i++ {
        fmt.Println(rand.Intn(50))
    }

    fmt.Println()
    for i := 0; i < 50; i++ {
        fmt.Println(rand.Float64())
    }

    fmt.Println()
    fmt.Println(rand.Float64() * 5)

    fmt.Println()
    s1 := rand.NewSource(42)
    r1 := rand.New(s1)

    fmt.Println("this is from s1:", r1.Intn(100))

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)

    fmt.Println("this is from s2:", r2.Intn(100))
}
