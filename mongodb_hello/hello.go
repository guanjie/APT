// test.go file for learning purpose
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	// 	"labix.org/v2/mgo/bson"
	// 	"strconv"
	"time"
)

type (
	Todo struct {
		Task      string    `bson:"task"`
		Created   time.Time `bson:"created"`
		Updated   time.Time `bson:"updated"`
		Completed time.Time `bson:"completed"`
	}
)

func main() {
	// declare 2 todos
	todo1 := Todo{Task: "1st test", Created: time.Now()}
	todo2 := Todo{Task: "2nd test", Created: time.Now()}

	// Open mongoSession
	fmt.Println("Awesome Eric!")
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	collection := mongoSession.DB("advent").C("wishes")
	collection.Insert(todo1)
	collection.Insert(todo2)

}
