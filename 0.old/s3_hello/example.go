package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.xiaonei.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println()

}
