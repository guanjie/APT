// temp.go file for learning purpose.
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", baiduhandler)
	http.ListenAndServe(":8080", nil)
}

func baiduhandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	w.Write(body)
}
