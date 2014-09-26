// auth_advanced.go file for learning purpose.
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
	m.Use(auth.BasicFunc(func(username, password string) bool { // Do auth.SecureCompare(word1, word1_compared)
		return auth.SecureCompare(username, "eric") && auth.SecureCompare(password, "isme")
	}))

	m.Get("/", func(w http.ResponseWriter) {
		fmt.Fprint(w, "Now we have this user correctly signed in")
	})

	m.Run()
}
