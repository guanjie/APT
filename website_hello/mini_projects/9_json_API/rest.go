// rest.go file for Json file manipulating and learning purpose.

package main

import (
	"encoding/json"
	"fmt"
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

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		log.Fatalf("getJsonResponse -> %v", err)
	}
	fmt.Fprintf(w, string(response))
}

func main() {
	fmt.Println("Awesome Eric!")

	http.HandleFunc("/", serveRest)
	http.ListenAndServe(":8080", nil)
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apple"] = 25
	fruits["Orange"] = 30

	veggies := make(map[string]int)
	veggies["Carrots"] = 21
	veggies["Tomatos"] = 40

	// Struct can directly be transformed into Json type by json.Marshal
	d := Data{fruits, veggies}
	p := Payload{d}

	return json.MarshalIndent(p, "", "   ")
}
