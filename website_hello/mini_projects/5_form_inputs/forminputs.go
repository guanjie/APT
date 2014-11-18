// form_inputs.go file for learning purpose.
// **EXTRA** Mongo Update, not replace

package main

import (
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	// 	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"strings"
)

var data *mgo.Collection

const (
	Address = ":8080"
)

type Login struct {
	Username []string `bson:"username,omitempty"`
	Password []string `bson:"password,omitempty"`
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("mgo.Dial to localhost -> %v", err)
	}
	return session
}

// General console and server info back
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // Parse url parameters passed.
	fmt.Println(r.Form) // print imformation on server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) // ???
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("value", strings.Join(v, "")) // ???
	}
	fmt.Fprintf(w, "Hello Eric!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) // get request method
	// Upon getting into the page it's a GET method.
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/login.gtpl")
		// This will then show to w.
		t.Execute(w, nil)
	} else { // ??? hence it's POST in here
		r.ParseForm()
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		// XXX Do this the update way, now it replaces
		data.Insert(Login{Username: r.Form["username"], Password: r.Form["password"]})
	}
}

func main() {
	// Get data - *mgo.Collection
	session := getSession()
	defer session.Close()
	data = session.DB("guanjie").C("login")

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(Address, nil)
	if err != nil {
		log.Fatalf("http.ListenAndServe -> %v", err)
	}
}
