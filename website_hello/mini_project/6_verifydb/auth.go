// auth.go file for learning purpose.
// **1** Adding account and password to our local mongodb.
// **2** Send the account and password back to the sender using Json format
// **3** Curl/http get the inforamtion I just returned

// **4** Go into the MongoDB pop out account == "human"
// **5** Put the additional tool to get the Mongodb check into tools folder

package main

import (
	// 	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

const databaseName string = "haoyun"
const collectionName string = "users"

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

func Mongo_connect_localhost() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func Mongo_insert_user(session *mgo.Session, account, password string) {
	collection := Mongo_get_collection(session, databaseName, collectionName)
	collection.Insert(&User{Account: account, Password: password, Created: time.Now()})
}

func Mongo_get_collection(session *mgo.Session, databaseName, collectionName string) *mgo.Collection {
	return session.DB(databaseName).C(collectionName)
}

func writeJson(w http.ResponseWriter, v interface{}) {
	// avoid json vulnerabilities, always wrap v in an object literal
	doc := map[string]interface{}{"d": v}

	if data, err := json.Marshal(doc); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func main() {
	m := martini.Classic()

	mongoSession := Mongo_connect_localhost()
	defer mongoSession.Close()
	m.Use(auth.BasicFunc(func(account, password string) bool {
		Mongo_insert_user(mongoSession, account, password) // **1**
		return account == "human" && password == "isme"
	}))

	m.Get("/", func(w http.ResponseWriter) {

	})

	m.Run()
}
