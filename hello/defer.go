// defer.go file, for learning purposes

package main

import (
    "fmt"
    "os"
)

func main() {
    // f is the *os.File that we defer close.
    // QUESTION can not pass in /tmp/tmp.txt, will not execute
    f := createFile("tmp.txt")
    // closeFile the *os.File f
    defer closeFile(f)
    writeFile(f)
}

// Creating file with the file name input, return the *os.File f
func createFile(p string) *os.File {
    // Print out that we are creating file
    fmt.Println("Creating file")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    // Write into *os.File f
    fmt.Fprintln(f, "Print into file\n and second line")
    fmt.Println("end of the printing")
}

func closeFile(f *os.File) {
    fmt.Println("Closing file")
    // *os.File f can directly be closed
    f.Close()
}
