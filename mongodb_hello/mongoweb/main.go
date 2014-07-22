// mongoweb main.go file for learning purpose.
package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
)

type Wish struct {
	Name        string `form: "name"`
	Description string `form:"description"`
}

func DB() martini.Handler {
	mongoSession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		// XXX session clone why?
		s := mongoSession.Clone()
		c.Map(s.DB("advent"))
		defer s.Close()
		c.Next()
	}
}

func GetAll(db *mgo.Database) []Wish {
	var wishlist []Wish
	// write all data into the wishlist Wish slice
	err := db.C("wishes").Find(nil).All(&wishlist)
	if err != nil {
		panic(err)
	}
	return wishlist
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	// After DB() we can use db as specific *database from context
	m.Get("/wishes", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "list", GetAll(db))
	})

	m.Post("/wishes", binding.Form(Wish{}), func(wish Wish, r render.Render, db *mgo.Database) {
		db.C("wishes").Insert(wish)
		r.HTML(200, "list", GetAll(db))
	})

	m.Run()
}
