package main

import "fmt"
import "strings"

func isAnagram(str1, str2 string) bool {
    if len(str1) == len(str2) {
        if strings.Contains(str1+str1, str2) {
            return true
        }
    }
    return false
}

func main() {
    str1 := "abcdefgh"
    str2 := "efghabcd"

    fmt.Println("Two strings str1 and str2 are anagrams?", isAnagram(str1, str2))
}
