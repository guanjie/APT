// test.go file for learning purpose
package main

import (
	"fmt"
	// 	"github.com/kr/pretty"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	// 	"strconv"
	"log"
	"time"
)

type (
	// We can do omitempty
	Todo struct {
		Task      string    `bson:"task,omitempty"`
		Created   time.Time `bson:"created,omitempty"`
		Updated   time.Time `bson:"updated,omitempty"`
		Completed time.Time `bson:"completed,omitempty"`
	}
)

func main() {
	// declare 2 todos
	// 	todo1 := Todo{Task: "1st test", Created: time.Now()}
	// 	todo2 := Todo{Task: "2nd test", Created: time.Now()}

	// Open mongoSession
	fmt.Println("Awesome Eric!")
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	collection := mongoSession.DB("advent").C("wishes")
	// 	collection.Insert(todo1)
	// 	collection.Insert(todo2)

	// Test 1: Adding a [string]interface{} map
	// We can also insert the map data that we pre-assumbled.
	// 	m := map[string]interface{}{
	// 		"name": "godep",
	// 		"tags": []string{"tool", "dependency"},
	// 		"contact": bson.M{
	// 			"name":  "Guanjie He",
	// 			"email": "humancool@gmail.com",
	// 		},
	// 	}
	// 	err = collection.Insert(m)
	if err != nil {
		log.Fatalf("collection insert -> %v", err)
	}
	fmt.Println("Okay!")

	// Test 2: Changing the information of a colleciton that I already got.
	// Set the additional data using "$set" inside the new M map[string]interface{}, tried Upsert, Update and UpdateAll
	type M map[string]interface{}
	change := M{"$set": Todo{Updated: time.Now()}}
	collection.Upsert(bson.M{"nbbbb": "nbme"}, change)
}
