// test.go file for learning purpose
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"time"
)

type (
	Todo struct {
		Task      string    `bson:"t"`
		Created   time.Time `bson:"c"`
		Updated   time.Time `bson:"u,omitempty"`
		Completed time.Time `bson:"cp,omitempty"`
	}
)

func main() {
	var todo1 = Todo{
		Task:    "Demo mgo",
		Created: time.Now(),
	}
	var todo2 = Todo{
		Task:    "this is cool",
		Created: time.Now(),
	}

	fmt.Println("Awesome Eric!")
	// Open mongoSession
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	// Connect to DB using mongoSession
	db := mongoSession.DB("names")
	collection := db.C("mynames")

	fmt.Println(collection.Count())
	fmt.Println(collection.FullName)

	collection.Insert(&todo1)
	collection.Insert(&todo2)

	fmt.Println(collection.Count())
	fmt.Println(collection.FullName)

}
