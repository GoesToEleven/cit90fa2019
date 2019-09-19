package main

import (
	"fmt"
	"net/http"
)

const key = "73626b1f"

func main() {
	r, err := http.Get(fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", key, "Gladiator"))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Printf("%T\n", r)
	fmt.Println(r)
}
