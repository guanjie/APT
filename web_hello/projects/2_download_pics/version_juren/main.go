// main.go file to download the pics from online XXX
// **1** Http get a first png file from online
// **2** put the http get file into laptop on Desktop
// **3** Download the file to Desktop folder
// **4** Write a function to get all the related links
// **5** Use the links to mass download the files
// **6** Goroutine to download all the pics at the same time
// **EXTRA** Ping and check all the numbers and get the most related links
// **EXTRA** Create folder and save files on desktop
// **EXTRA** Get all the folders ready and download all the manga pages for onepiece

// **EXTRA** Refactor the download_to_desktop function, inside function it creates a folder every time.

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

	// Write file from data to file path
	splitup := strings.Split(url, "/")
	foldername := splitup[len(splitup)-2]
	filename := splitup[len(splitup)-1]
	// XXX hardcoded path XXX
	if err := os.MkdirAll("/Users/guanjiehe/Desktop/juren/"+foldername, 0766); err != nil {
		log.Fatalf("os.MkdirAll -> %v", err)
	}
	if err := ioutil.WriteFile("/Users/guanjiehe/Desktop/juren/"+foldername+"/"+filename, data, 0766); err != nil {
		log.Fatalf("ioutil.WriteFile -> %v", err)
	}
}

func all_related_links(url string) (urls []string) {
	// XXX NEED TO MODIFY XXX for now hard coded numbers: 1 to 50
	for i := 1; i <= 50; i++ {
		urls = append(urls, url+strconv.Itoa(i)+".jpg")
	}
	return
}

func mass_download_to_desktop(urls []string) {
	for _, url := range urls {
		// Check if there is a 404, if there is, break and return
		resp, err := http.Get(url)
		if resp.StatusCode == 404 {
			return
		}
		if err != nil {
			log.Fatalf("http.Get in mass download func -> %v", err)
		}
		resp.Body.Close()

		go download_to_desktop(url)
	}
}

func main() {
	fmt.Println("Awesome Eric!")
	baseUrl := "http://2.p.mpcdn.net/20977/"
	const startIndex = 349458
	const endIndex = 349500

	for i := startIndex; i <= endIndex; i++ {
		urls := all_related_links(baseUrl + strconv.Itoa(i) + "/")
		mass_download_to_desktop(urls)
	}

	var input string
	fmt.Scanln(&input)
}
