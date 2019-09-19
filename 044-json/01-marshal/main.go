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
}
