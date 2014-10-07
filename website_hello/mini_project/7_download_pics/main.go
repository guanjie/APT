//main.go file to download the pics from online XXX
// **1** Http get a first png file from online
// **2** put the http get file into laptop on Desktop
// **3** Download the file to Desktop folder
// **4** Write a function to get all the related links
// **5** Use the links to mass download the files

// **6** Goroutine to download all the pics at the same time
// **EXTRA** Ping and check all the numbers and get the most related links

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func download_to_desktop(url string) {
	// Get response from the jpg link
	resp, err := http.Get(url)
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
	filename := splitup[len(splitup)-1]
	if err := ioutil.WriteFile("/Users/guanjiehe/Desktop/"+filename, data, 0666); err != nil {
		log.Fatalf("ioutil.WriteFile -> %v", err)
	}
}

func all_related_links(url string) (urls []string) {
	// XXX NEED TO MODIFY XXX for now hard coded numbers: 1 to 32
	for i := 1; i <= 32; i++ {
		urls = append(urls, url+strconv.Itoa(i)+".png")
	}
	return
}

func mass_download_to_desktop(urls []string) {
	for _, url := range urls {
		download_to_desktop(url)
	}
}

func main() {
	fmt.Println("Awesome Eric!")
	baseUrl := "http://2.p.mpcdn.net/10799/164439/"

	urls := all_related_links(baseUrl)
	mass_download_to_desktop(urls)
}
