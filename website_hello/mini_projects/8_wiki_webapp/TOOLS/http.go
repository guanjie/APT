// http.go file for learning purpose. Has the http.HandleFunc function and ListenAndServe
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the responseWriter: %s", r.URL.Path[1:])
}
