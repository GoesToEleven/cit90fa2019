package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type movie struct {
	Title    string `json:"Title"`
	Director string `json:"Director"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
}

const key = "73626b1f"

func main() {
	r, err := http.Get(fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", key, "Gladiator"))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

	m := movie{}
	err = json.Unmarshal(bs, &m)
	fmt.Println(m.Title)
	fmt.Println(m.Director)
	fmt.Println(m.Actors)
	fmt.Println(m.Plot)
	fmt.Println(m.Ratings)
}
