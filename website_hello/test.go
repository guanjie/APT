// test.go file for basic template and laerning purpose.

package main

import (
	// 	"fmt"
	// 	"github.com/codegangsta/martini"
	"io/ioutil"
	"log"
	"net/http"
	// 	"os"
)

func main() {
	// 	// 	test1 - martini
	// 	m := martini.Classic()
	//
	// 	m.Get("/", func(w http.ResponseWriter) {
	// 		resp, _ := http.Get("http://2.p.mpcdn.net/10799/164300/40.jpg")
	// 		// 		status := resp.Status
	// 		statusCode := resp.StatusCode
	// 		// 		fmt.Fprint(w, status)
	// 		fmt.Fprint(w, statusCode)
	// 	})
	// 	m.Run()

	// test2 - Check os.MkdirAll
	// 	err := os.MkdirAll("/Users/guanjiehe/Desktop/try", 0777)
	// 	if err != nil {
	// 		log.Fatalf("os.MkdirAll -> %v", err)
	// 	}

	// test3 - actually download the flv file...Be aware this is porn XXX
	var url = "http://im.c6881952.c59b2a8.cdn2c.videos2.youjizz.com/c/1/c11fa9963ac4218933ea87c8091b71c21413528902-640-480-1800-h264.flv?rs=350&ri=1400&s=1413687934&e=1413860734&h=0c048d8afbe5c431a036393c61d21765"
	download_to_desktop(url)
}

func download_to_desktop(url string) {
	// Get response from the jpg link
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		log.Fatal("404 code error")
		return
	}
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	defer resp.Body.Close()

	// Get the data by using ioutil.ReadAll(resp.Body)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// XXX hardcoded path XXX
	if err := ioutil.WriteFile("/Users/guanjiehe/Desktop/checkit.flv", data, 0766); err != nil {
		log.Fatalf("ioutil.WriteFile -> %v", err)
	}
}
