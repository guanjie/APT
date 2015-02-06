// hashes.go file for learning purpose. There are three different encryption algorithms:

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Awesome Eric!")

	h := sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, "The second string in here.")
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, "Tried the third one in here.")
	fmt.Printf("%x\n", h.Sum(nil))
}
