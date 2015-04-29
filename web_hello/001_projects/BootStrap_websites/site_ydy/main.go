// martini.go for learning purpose
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func coolhandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("static/ydy.html")
	fmt.Fprintf(w, string(body))
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)

}
