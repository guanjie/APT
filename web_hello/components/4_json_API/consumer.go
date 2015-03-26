// consumer.go file to unmarshal the data into our Payload struct data type. For learning
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func main() {
	fmt.Println("Awesome Eric!")

	url := "http://localhost:8080"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	defer res.Body.Close()

	// This is []byte type
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	var p Payload

	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatalf("json.Unmarshal -> %v", err)
	}

	fmt.Println("After getting data we have:", p)
}
