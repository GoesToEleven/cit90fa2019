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
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		log.Panicf("error marshalling data: %s", err)
	}

	// fmt.Println(string(bs))

	fmt.Fprint(w, string(bs))
}
