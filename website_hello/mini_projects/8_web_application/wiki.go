// wiki.go file for learning purpose. Check the from golang.org/doc/articles/wiki/.
package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

// To cache the templates first
var templates = template.Must(template.ParseFiles("./templates/edit.html", "./templates/view.html"))
var validPath = regexp.MustCompile("^/(edit|view|save)/([0-9a-zA-Z]+)$")

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("./page_data/"+filename, p.Body, 0766) // XXX ioutil deals with []byte type
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("./page_data/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	match := validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		log.Fatalf("Not found validPath")
		return "", errors.New("Invalid Title")
	}
	return match[2], nil
}

func renderTemplate(w http.ResponseWriter, tmplName string, p *Page) {
	err := templates.ExecuteTemplate(w, tmplName+".html", p)
	if err != nil {
		log.Fatalf("renderTempalte parsefiles error -> %v", err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		log.Fatalf(" viewHandler getTitle-> %v", err)
	}
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		log.Fatalf("editHandler getTitle -> %v", err)
	}
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		log.Fatalf("saveHandler getTitle -> %v", err)
	}

	// XXX hard part, request is a POST request of the form information
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	if err := page.save(); err != nil {
		log.Fatalf("page.save() -> %v", err)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	fmt.Println("Awesome Eric!")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
