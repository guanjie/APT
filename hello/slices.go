// slices.go file for learning purposes

package main

import (
    "fmt"
)

var p = fmt.Println

func main() {

    // TRICK, s := make([]string, 3), need to specify the length of the slices
    s := make([]string, 4)

    s[0] = "First string"
    s[1] = "Second string"
    s = append(s, "fourth")
    s = append(s, "5", "6")

    c := make([]string, len(s))
    copy(c, s)
    s[0] = "changed"
    p(s)
    p(c)

    l := s[0:4]
    p(l)

}
