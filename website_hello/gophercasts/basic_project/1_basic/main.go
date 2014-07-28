// main.go file for basic learning purpose.
package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/generate", func(r *http.Request) string {
		body := r.FormValue("body")
		return body
	})

	m.Run()
}
