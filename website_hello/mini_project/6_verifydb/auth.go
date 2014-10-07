// auth.go file for learning purpose.
// **1** Adding account and password to our local mongodb.
// **2** Send the account and password back to the sender using Json format
// **3** Curl/http get the inforamtion I just returned
// **4** Find out account == "human", put into interface, print it out to w

// **5** Put the additional tool,and http://denvergophers.com/2013-04/mgo.article into a tools file.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"strconv"
	"time"
)

const databaseName string = "haoyun"
const collectionName string = "users"

type (
	UserRepo struct {
		Collection *mgo.Collection
	}

	Users []User
	User  struct {
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

func (r UserRepo) All() (users Users, err error) {
	err = r.Collection.Find(bson.M{}).All(&users)
	return
}

func main() {
	fmt.Println("Eric is awesome")
	m := martini.Classic()

	mongoSession := Mongo_connect_localhost()
	defer mongoSession.Close()
	m.Use(auth.BasicFunc(func(account, password string) bool {
		Mongo_insert_user(mongoSession, account, password) // **1**
		return account == "eric" && password == "isme"
	}))

	m.Get("/", func(w http.ResponseWriter) {
		var (
			users []User // Import bunch of users
			err   error
		)

		myCollection := Mongo_get_collection(mongoSession, databaseName, collectionName)
		var repo UserRepo = UserRepo{
			Collection: myCollection,
		}

		if users, err = repo.All(); err != nil {
			panic(err)
		}
		writeJson(w, users)
	})

	m.Run()
}
