// test.go file for basic template and laerning purpose.

package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		res, _ := http.Get("http://www.baidu.com")
		status := res.Status
		return status
	})

	m.Run()
}
