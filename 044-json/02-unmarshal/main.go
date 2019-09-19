package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	p2 := person{}
	fmt.Println("p2 before unmarshal", p2)
	
	err = json.Unmarshal(bs, &p2)
	if err != nil {
		panic(err)
	}
	fmt.Println("p2 after unmarshal", p2)
	// json --> go
	// https://mholt.github.io/json-to-go/
}
