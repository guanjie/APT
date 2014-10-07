// auth_user.go file for learning purpose.
package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"net/http"
)

func main() {
	m := martini.Classic()
	// authenticate every request

	// m.Use() uses data and saved data inside directly
	m.Use(auth.BasicFunc(func(username, password string) bool {
		return username == "eric" && password == "isme"
	}))

	m.Get("/", func(w http.ResponseWriter, user auth.User) {
		fmt.Fprint(w, user)
	})

	m.Run()
}
