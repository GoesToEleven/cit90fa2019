package main

import (
	"io"
	"net/http"
	"strconv"
)
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "some string")
	xs := []string{"jenny", "james", "ian"}
	s := ""
	for i, v := range xs {
		s += "<h1>" + strconv.Itoa(i) + " - HELLO " + v + "</h1>"
	}
	io.WriteString(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
	`+s+`
	</body>
	</html>
	`)
}
