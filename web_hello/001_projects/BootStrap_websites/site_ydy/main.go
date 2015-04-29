// martini.go for learning purpose
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func coolhandler(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "cool man")
	body, _ := ioutil.ReadFile("static/ydy.html")
	fmt.Fprint(w, string(body))
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	// 	http.HandleFunc("/", coolhandler)
	http.ListenAndServe(":8080", nil)

}
