// test.go file for learning purpose
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	// 	"labix.org/v2/mgo/bson"
	"time"
)

type (
	OfflineService struct {
		Servicename string `bson:"servicename"` // Service name for the service
		Price       int32  `bson:"price"`
		Desc        string `bson:"description"`
		Category    string `bson:"category"`
		Address     string `bson:"address"`
		Comments_id string `bson:"comments_id"`
		Banner      string `bson:"banner_link"`
		Picture     string `bson:"picture_link"`
		Merchant    string `bson:"merchant"`

		Created time.Time `bson:"created"`
		Updated time.Time `bson:"updated"`
	}
)

func main() {
	// 	user1 := User{Task: "1st test", Created: time.Now()}
	// 	user2 := User{Task: "2nd test", Created: time.Now()}

	service1 := OfflineService{Servicename: "strip dance", Price: 120, Desc: "You will feel like king afterwards.", Category: "Entertainment", Address: "asdfasdfasdf asdfasdf", Merchant: "Beijing Hot.", Created: time.Now()}
	service2 := OfflineService{Servicename: "strip dance2", Price: 20000, Desc: "You will love this", Category: "Entertainment", Address: "Your home baby", Merchant: "Beijing Hot.", Created: time.Now()}

	// Open mongoSession
	fmt.Println("Awesome Eric!")
	mongoSession, err := mgo.Dial("localhost") // Adding the entries into localhost from server
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	collection := mongoSession.DB("haoyun").C("offlineservices")
	collection.Insert(service1)
	collection.Insert(service2)

}
