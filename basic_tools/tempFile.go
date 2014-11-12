//temp.go file for learning purpose.
package main

import (
	"fmt"
	"io/ioutil"
	// 	"log"
	"os"
)

func main() {
	fmt.Println("Awesome Eric!")

	// ioutil.TempFile(), we can use the os.TempDir() to get the tempDir() up.
	file, err := ioutil.TempFile("./", "prefix")
	if err != nil {
		log.Fatalf("ioutil.TempFile -> %v", err)
	}
	defer os.Remove(file.Name())

}
