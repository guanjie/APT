package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println()

}
