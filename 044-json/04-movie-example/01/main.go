package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
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
