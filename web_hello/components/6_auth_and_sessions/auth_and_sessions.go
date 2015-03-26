// auth_and_sessions.go file to handle user registration and anthentication
// Able to check Login credentials and pass User back to backend
// **EXTRA** Sessions usage and coockies usage, using gorilla - http://shadynasty.biz/blog/2012/09/05/auth-and-sessions/

package main

import (
	"code.google.com/p/go.crypto/bcrypt" // bcrypt
	"code.google.com/p/gorilla/sessions"
	"fmt"
	"html/template"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty` // ???
	Username string
	Password []byte // note: uses bcrypt hashes to set the password.

	Posts int
}

// SetPassword takes a plaintext password and hashes it with bcrypt and set the password filed to the hash
func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("bcrypt.GenerateFromPassword -> %v", err)
	}
	u.Password = hpass
}

// Login validates and returns a user, if user exists in database
func Login(ctx *Context, username, password string) (u *User, err error) {
	err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
	if err != nil {
		// if error finding the username, return
		return
	}

	// if there is NO error, check password
	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		u = nil // if there is NO error before, but passwords don't match
		return
	}
	return
}

// xxx can be better
var templates = template.ParseFiles("templates/_base.html", "templates/login.html")

func loginForm(w http.ResponseWriter, r *http.Request, ctx *Context) (err error) {
	err = templates.Execute(w, nil)
	return
}

func login(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	username, password := req.FormValue("username"), req.FormValue("password")

	// login the user
	user, err := Login(ctx, username, password)

	// if error, present the form again with error message

	// if NO error, user login is valid -> What to do? Sessions
	// XXX
}

func main() {
	fmt.Println("Awesome Eric!")
}
