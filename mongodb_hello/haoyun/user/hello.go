// test.go file for learning purpose
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	// 	"labix.org/v2/mgo/bson"
	"time"
)

type (
	User struct {
		Account  string `bson:"account"` // Virtual name used in app
		Password string `bson:"password"`
		Name     string `bson:"name"`
		Avatar   string `bson:"avatar_link"`
		Week     int    `bson:"week"`
		Weekborn int    `bson:"weekborn`

		Online_service_id  string `bson:"online_service_id"`
		Offline_service_id string `bson:"offline_service_id"`

		Created time.Time `bson:"created"`
		Updated time.Time `bson:"updated"`
	}
)

func main() {
	// 	user1 := User{Task: "1st test", Created: time.Now()}
	// 	user2 := User{Task: "2nd test", Created: time.Now()}

	user1 := User{Account: "humancool", Password: "PASSWORDISINHERE", Online_service_id: "asdfasdfwekljsadf", Name: "冠杰何", Week: 10, Created: time.Now()}
	user2 := User{Account: "Harvey", Name: "冠豪", Week: 20, Created: time.Now()}

	// Open mongoSession
	fmt.Println("Awesome Eric!")
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()

	collection := mongoSession.DB("haoyun").C("users")
	collection.Insert(user1)
	collection.Insert(user2)

}
