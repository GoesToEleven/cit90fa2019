package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO WORLD")
}

func bar(w http.ResponseWriter, r *http.Request) {
	name := "JAMES BOND"
	s := "one thing" + "another thing"
	fmt.Fprint(w, `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>HELLO` + name + `</title>
    <style>
        h1 {
            color: red;
        }
    </style>
</head>
<body>

    <h1>WELCOME TO THE SHOW` + name + `</h1>
    <p>Lorem, ipsum dolor sit amet consectetur adipisicing elit. Fugiat molestiae dignissimos nulla maiores accusamus, doloribus tempora delectus dolorum quas consectetur sunt, aspernatur unde blanditiis reiciendis? Repudiandae adipisci enim voluptatibus magnam?</p>
	<h2>` + s + `</h2>
	<a href="/">GO HOME</a>
</body>
</html>
	`)
}
