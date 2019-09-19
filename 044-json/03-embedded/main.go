package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Favs        []string
	Drinks      map[int]string
}

type secretAgent struct {
	person
	Ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			First:  "James",
			Last:   "Bond",
			Favs:   []string{"behind you", "nobody lives forever", "never say never"},
			Drinks: map[int]string{1: "martini", 2: "gin"},
		},
		Ltk: true,
	}
	fmt.Println("HERE", sa1)
	bs, err := json.MarshalIndent(sa1, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	sa2 := secretAgent{}
	fmt.Printf("sa2 before unmarshal %+v\n", sa2)

	err = json.Unmarshal(bs, &sa2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sa2 after unmarshal %+v\n", sa2)
	// json --> go
	// https://mholt.github.io/json-to-go/
}
