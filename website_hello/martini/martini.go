// martini.go for learning purpose

package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Post("/generate", func(r *http.Request) string {
		return r.FormValue("body")
	})

	m.Run()
}
