package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action="/process" method="GET">
        <input type="text" name="first" placeholder="FIRST NAME">
        <input type="submit">
    </form>
</body>
</html>
	`)
}

func process(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")

	w.Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>"+f+"</h1>")
}
