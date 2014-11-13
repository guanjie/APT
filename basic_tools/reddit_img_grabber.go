// concurrently download images from any subreddit
package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type JsonResponse struct {
	Kind string
	Data DataType
}

type DataType struct {
	Children []ChildrenType
}

type ChildrenType struct {
	Data DataType2
}

type DataType2 struct {
	Url string
}

type RedditImage struct {
	url       string
	name      string
	http_resp *http.Response
}

func main() {
	fmt.Println("Awesome Eric!")

	var subreddit, directory string
	fmt.Print("Subreddit: ")
	fmt.Scanln(&subreddit)
	fmt.Print("Path where to save: ")
	fmt.Scanln(&directory)
	grab(subreddit, directory)
}

func grab(sub string, dir string) {
	ch := make(chan RedditImage)
	urls := geturls(sub)

	// INTERESTING, go fetchimage put http_resp into redditimg, then put it in ch
	for _, url := range urls {
		reddit := RedditImage{url: url, name: urltoname(url)}
		go fetchimage(reddit, ch)
	}

	for {
		select {
		case redditimg := <-ch:
			saveimage(redditimg, dir)
		}
	}
}

func geturls(subreddit string) []string {
	// get http response
	response, err := http.Get("http://reddit.com/r/" + subreddit + ".json")
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	defer response.Body.Close()

	// get response.Body as contents
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// get JsonResponse using unmarshal
	var resp JsonResponse
	if err := json.Unmarshal(contents, &resp); err != nil {
		log.Fatalf("json.Unmarshal -> %v", err)
	}

	// get urls []string from children type
	children := resp.Data.Children
	var urls []string
	for _, val := range children {
		urls = append(urls, val.Data.Url)
	}
	return urls
}

func urltoname(url string) string {
	split := strings.Split(url, "/")
	lenofsplit := len(split)
	indexoflast := lenofsplit - 1
	namewithext := split[indexoflast]
	split2 := strings.Split(namewithext, ".")
	name := split2[0]
	return name
}

func fetchimage(redditimg RedditImage, ch chan RedditImage) {
	// complete redditimg by adding http_resp
	resp, err := http.Get(redditimg.url)
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	redditimg.http_resp = resp

	// put resp into ch
	ch <- redditimg
}

func saveimage(redditimg RedditImage, dir string) {
	name := redditimg.name

	// decode image
	img, _, _ := image.Decode(redditimg.http_resp.Body)
	if img != nil {
		file, _ := os.Create(dir + name + ".png")
		err := png.Encode(file, img)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
}
