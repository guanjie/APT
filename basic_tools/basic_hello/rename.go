package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Awesome Eric!")
	f, _ := os.Open("/Users/guanjiehe/Desktop/KKRheiheihei.pdf")
	os.Rename(f.Name(), "/Users/guanjiehe/Desktop/KKRhe.pdf")
}
