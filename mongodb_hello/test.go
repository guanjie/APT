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
	// declare 2 todos
	// 	var todo1 = Todo{
	// 		Task:    "Demo mgo",
	// 		Created: time.Now(),
	// 	}
	// 	var todo2 = Todo{
	// 		Task:    "this is cool",
	// 		Created: time.Now(),
	// 	}

	// Open mongoSession
	fmt.Println("Awesome Eric!")
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	// Connect to DB using mongoSession
	db := mongoSession.DB("testing")
	collection := db.C("eric")
	fmt.Println(collection.Count())
	fmt.Println(collection.FullName)
	fmt.Println(collection.Database)

	// Insert 2 todos, by inserting the address. ATTN: & or no &
	// 	collection.Insert(todo1)
	// 	collection.Insert(todo2)
	//
	// 	fmt.Println(collection.Count())
	// 	fmt.Println(collection.FullName)

}
