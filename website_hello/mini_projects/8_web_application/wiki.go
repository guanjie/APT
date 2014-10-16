// wiki.go file for learning purpose. Check the from golang.org/doc/articles/wiki/.
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0766)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile -> %v", err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		log.Fatalf(" loadPage from Viewhandler-> %v", err)
		page = &Page{Title: "empty"}
	}

	t, _ := template.ParseFiles("view.html")
	t.Execute(w, page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		log.Fatalf("LoadPage -> %v", err)
		page = &Page{Title: title}
	}

	t, _ := template.ParseFiles("./templates/edit.html")
	t.Execute(w, page)
}

func main() {
	fmt.Println("Awesome Eric!")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// 	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
