// main.go file to download the pics from online XXX
// **1** Http get a first png file from online
// **2** put the http get file into laptop on Desktop
// **3** Download the file to Desktop folder
// **4** Write a function to get all the related links
// **5** Use the links to mass download the files
// **6** Goroutine to download all the pics at the same time
// **EXTRA** Ping and check all the numbers and get the most related links
// **EXTRA** Create folder and save files on desktop
// **EXTRA** Get all the folders ready and download all the manga pages for naruto

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input info
	ch := make(chan []string, 10)
	baseUrl := "http://2.p.mpcdn.net/20606/"
	const startIndex = 338065
	const endIndex = 338600

	// get all urls and mass download them
	for i := startIndex; i <= endIndex; i++ {
		// get episode urls
		urls := geturls(baseUrl + strconv.Itoa(i) + "/")
		// HINT: Use time.sleep/ time.after
		go func() {
			ch <- urls
		}()
	}

	for {
		select {
		case urls := <-ch:
			saveimages(urls)
		}
	}

	var input string
	fmt.Scanln(&input)
}

// get episode urls
func geturls(url string) (urls []string) {
	// XXX NEED TO MODIFY XXX for now hard coded numbers: 1 to 50
	for i := 1; i <= 50; i++ {
		urls = append(urls, url+strconv.Itoa(i)+".jpg")
	}
	return
}

func saveimages(urls []string) {
	for _, url := range urls {
		go download_to_desktop(url)
	}
}

func download_to_desktop(url string) {
	// Get response from the jpg link
	resp, _ := http.Get(url)
	if resp.StatusCode == 404 {
		// don't do anything
		fmt.Println("404 code error from url:", url)
	}
	defer resp.Body.Close()

	// Get the data by using ioutil.ReadAll(resp.Body)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// MkdirAll and WriteFile
	foldername, filename := urltonames(url)
	if err := os.MkdirAll("/Users/guanjiehe/Desktop/naruto/"+foldername, 0766); err != nil {
		log.Fatalf("os.MkdirAll -> %v", err)
	}
	if err := ioutil.WriteFile("/Users/guanjiehe/Desktop/naruto/"+foldername+"/"+filename, data, 0766); err != nil {
		log.Fatalf("ioutil.WriteFile -> %v", err)
	}
}

func urltonames(url string) (string, string) {
	// Write file from data to file path
	splitup := strings.Split(url, "/")
	foldername := splitup[len(splitup)-2]
	filename := splitup[len(splitup)-1]
	return foldername, filename
}
