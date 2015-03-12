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
	m.Use(auth.Basic("eric", "awesome")) // This is the middleware

	m.Get("/", func(w http.ResponseWriter) { // Direct insert info into response
		fmt.Fprint(w, "Eric is awesome")
	})

	m.Run()
}
