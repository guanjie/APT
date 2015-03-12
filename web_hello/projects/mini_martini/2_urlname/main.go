// martini.go for learning purpose
package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
)

func main() {
	m := martini.Classic()

	m.Get("/", func(r *http.Request) string {
		return "Hello world"
	})

	m.Get("/number/:num", func(params martini.Params) string {
		num_str := params["num"]
		num, _ := strconv.Atoi(num_str)
		num = powerto3(num)
		return "Hello the number " + params["num"] + " to the power 3 is: " + strconv.Itoa(num)
	})

	m.Get("/hello/:name/:verb" /* Auth,*/, func(params martini.Params) string {
		return "Hello " + params["name"] + " and you are " + params["verb"]
	})

	m.Run()
}

func powerto3(num int) int {
	return num * num * num
}

/*
func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "secret123" {
		http.Error(res, "Nope", 401)
	}
}
*/
