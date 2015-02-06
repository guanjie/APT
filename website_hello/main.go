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

	// 算法开始
	m.Get("/number/:num", func(params martini.Params) string {
		num, _ := strconv.Atoi(params["num"])
		num = powerto3(num)
		return "Hello the number " + params["num"] + " to the power 3 is: " + strconv.Itoa(num)
	})

	m.Run()
}

func powerto3(num int) int {
	return num * num * num
}
