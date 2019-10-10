package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}
func bar(w http.ResponseWriter, r *http.Request) {
	people2 := []person{}
	err := json.NewDecoder(r.Body).Decode(&people2)
	if err != nil {
		http.Error(w, "couldn't decode the json - bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, people2)
	/*
		https://curlbuilder.com/

		curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"},{"First":"James"}]' 'http://localhost:8080/decode'
	*/
	fmt.Println("curl ran")
	fmt.Println("curl ran")
	fmt.Println("curl ran")
	fmt.Println("here's the data")
	fmt.Println(people2)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	people := []person{p1, p2}

	// bs, err := json.Marshal(people)
	// if err != nil {
	// 	log.Panicf("error marshalling data: %s", err)
	// }

	// // fmt.Println(string(bs))

	// fmt.Fprint(w, string(bs))

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Panicf("couldn't send encoded data %s", err)
	}
}
