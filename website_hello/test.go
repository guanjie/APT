// test.go file for basic template and laerning purpose.

package main

import (
	// 	"fmt"
	// 	"github.com/codegangsta/martini"
	// 	"net/http"
	"log"
	"os"
)

func main() {
	// 	// 	test1 - martini
	// 	m := martini.Classic()
	//
	// 	m.Get("/", func(w http.ResponseWriter) {
	// 		resp, _ := http.Get("http://2.p.mpcdn.net/10799/164300/40.jpg")
	// 		// 		status := resp.Status
	// 		statusCode := resp.StatusCode
	// 		// 		fmt.Fprint(w, status)
	// 		fmt.Fprint(w, statusCode)
	// 	})
	// 	m.Run()

	// test2 - Check os.MkdirAll
	err := os.MkdirAll("/Users/guanjiehe/Desktop/try", 0777)
	if err != nil {
		log.Fatalf("os.MkdirAll -> %v", err)
	}
}
