// test.go file for basic template and laerning purpose.

package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func(w http.ResponseWriter) {
		res, _ := http.Get("http://www.baidu.com")
		status := res.Status
		fmt.Fprint(w, status)
	})

	m.Run()
}
