// martini.go for learning purpose
package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func(r *http.Request) string {
		return "Hello world"
	})

	m.Get("/hello/:name/:verb", Auth, func(params martini.Params) string {
		return "Hello " + params["name"] + " and you are " + params["verb"]
	})

	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "secret123" {
		http.Error(res, "Nope", 401)
	}
}
