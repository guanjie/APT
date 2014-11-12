// wiki.go file for learning purpose. Check the from golang.org/doc/articles/wiki/.
// Write the data into the local database, call the mgo.Dial to connect to localhost.
// 1. Check view() and see what need to be done, make the change
// 2. Save() function now saves to local host, see what could be done
// 3. loadpage() function see what need to be done.
// 4. Editpage() function see what need to be done.
// 5. savepage() function see what need to be done.
// 6. Get to pull off the set up the collection part from main()

// 7. Now make the change to make module support adding another field in the struct.

package main

import (
	"errors"
	"fmt"
	"html/template"
	// 	"io/ioutil"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string `bson:title,omitempty`
	Body  []byte `bson:body,omitempty`
}

var templates = template.Must(template.ParseFiles("./templates/edit.html", "./templates/view.html"))
var validPath = regexp.MustCompile("^/(edit|view|save)/([0-9a-zA-Z]+)$")
var pages *mgo.Collection

func getSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("mgo.Dial -> %v", err)
	}
	return session
}

// XXX Save to database
func (p *Page) save() error {
	// Save to database
	return pages.Insert(p)
}

func loadPage(title string) (*Page, error) {
	var page Page
	// Remember that we are fetching data from collection, use bson.M
	err := pages.Find(bson.M{"title": title}).One(&page)
	if err != nil {
		fmt.Printf("collection.Find().One() -> %v", err)
		return nil, err
	}
	return &page, nil
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

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		match := validPath.FindStringSubmatch(r.URL.Path)
		if match == nil {
			log.Fatalf("MakeHandler regexp match ", nil)
		}
		// XXX remember we don't return any valuat this step, wrapper class
		fn(w, r, match[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	if err := page.save(); err != nil {
		log.Fatalf("page.save() -> %v", err)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	session := getSession()
	defer session.Close()

	// In main the pages was declared up front
	pages = session.DB("wiki").C("pages")

	fmt.Println("Awesome Eric!")
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
