// auth.go file for learning purpose.

package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"net/http"
)

func main() {
	fmt.Println("Awesome Eric!")
	m := martini.Classic()
	m.Use(auth.Basic("username", "secretpassword"))

	m.Get("/", func(w http.ResponseWriter) {
		fmt.Fprint(w, "Eric is awesomeeeee")
	})

	m.Run()
}
