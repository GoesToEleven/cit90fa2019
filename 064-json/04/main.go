package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
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

	fmt.Println(string(bs))

	people2 := []person{}
	err = json.Unmarshal(bs, &people2)
	fmt.Println("HERE HERE HERE")
	fmt.Println(people2)
}
