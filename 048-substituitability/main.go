package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprint(f, "hello world first")
	fmt.Fprint(os.Stdout, "hello world second\n\n")

	http.HandleFunc("/", booga)
	http.ListenAndServe(":8080", nil)
}

func booga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world again")
}
