package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/tobs.jpeg", tob)
	http.ListenAndServe(":8080", nil)
}

func tob(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("tobs.jpeg")
	if err != nil {
		http.Error(w, "couldn't open file", http.StatusInternalServerError)
	}
	defer f.Close()

	io.Copy(w, f)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<img src="/tobs.jpeg" alt="cute dog">
	</body>
	</html>
	`)
}
