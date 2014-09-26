// auth.go file for learning purpose.
// **1** Adding account and password to our local mongodb.
// **2** Send the account and password back to the sender using Json format
// **3** Curl/http get the inforamtion I just returned

package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"labix.org/v2/mgo"
	"net/http"
	"time"
)

func main() {
	m := martini.Classic()
	m.Use(auth.BasicFunc(func(account, password string) bool {
		Mongo_insert_user(account, password) // **1**
		return account == "eric" && password == "isme"
	}))

	m.Get("/", func(w http.ResponseWriter, user auth.User) []byte {
		resultjson, err := json.Marshal(map[string]interface{}{"account": string(user)}) // **2**
		if err != nil {
			panic(err)
		}
		return resultjson
	})

	m.Run()
}

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

func Mongo_insert_user(account, password string) {
	user := User{Account: account, Password: password, Created: time.Now()}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("haoyun").C("users")
	c.Insert(user)
}
