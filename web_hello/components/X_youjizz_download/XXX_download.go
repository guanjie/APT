// test.go file for basic template and learning purpose. XXX
// XXX Need to add regexp function to fetch the .flv file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	// actually download the flv file...Be aware this is porn XXX
	var url2 = "http://www.megafilex.com/media/player/jw/player.swf?config=http://www.megafilex.com/media/player/jw/config_embed.php?vkey=7191"
	// 	var url2 = "http://im.c6881952.17110d1.cdn2c.videos2.youjizz.com/2/0/201540cf6703ae4a1ba9bd3def94bdcf1393303502-640-480-2140-h264.flv?rs=350&ri=1400&s=1413702543&e=1413875343&h=a3441212bb7f84e4c386caf2e6ea1058"
	// 	var url3 = "http://im.c6881952.ea329dc.cdn1c.videos2.youjizz.com/1/2/129d9fb88f8b6ff36c4a9ced0dd6a198.flv?rs=300&ri=1200&s=1413702545&e=1413875345&h=9cdb7f444372b060d77ef3063cf5623a"
	// 	urllist := []string{url2, url3}
	urllist := []string{url2}
	for _, url := range urllist {
		go download_to_desktop(url)
	}
	var input string
	fmt.Scanln(&input)
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

	filename := time.Now().String()
	// XXX hardcoded path XXX
	if err := ioutil.WriteFile("/Users/guanjiehe/Desktop/"+filename[20:26]+".flv", data, 0766); err != nil {
		log.Fatalf("ioutil.WriteFile -> %v", err)
	}
}
