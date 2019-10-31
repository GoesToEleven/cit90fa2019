package main

import (
	"io"
	"net/http"
)

type user struct {
	First    string
	Email    string
	Password string
}

// create variable to hold value of TYPE map[string]user
// DOES NOT INITIALIZE variable with VALUE in memory
// var db map[string]user

// create variable to hold value of TYPE map[string]user
// INITIALIZES variable with VALUE in memory
var db = map[string]user{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signupprocess", processSignUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginprocess", processLogin)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u, ok := db[c.Value]
	if !ok {
		u = user{}
	}

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

    <a href="/signup">SIGNUP</a>
    <a href="/login">LOGIN</a>
    <a href="/logout">LOGOUT</a>
	<h1>IF YOU ARE LOGGED IN YOU WILL SEE YOUR NAME:</h1>
	`+u.First+`  
</form>
</body>
</html>
`)
}

func signup(w http.ResponseWriter, r *http.Request) {
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
    <form action="/signupprocess" method="POST">
    <input type="text" name="fn" placeholder="FIRST NAME">
    <input type="text" name="em" placeholder="EMAIL">
    <input type="password" name="pw" placeholder="PASSWORD">
    <input type="submit">    
</form>
</body>
</html>
	`)
}

func processSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	f := r.FormValue("fn")
	if f == "" {
		http.Error(w, "first name cannot be empty", http.StatusBadRequest)
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
	}

	// NOTE:
	// never directly store a password
	// we will fix this later ....
	u := user{
		First:    f,
		Email:    e,
		Password: p,
	}

	db[e] = u

	c := &http.Cookie{
		Name:  "user-session",
		Value: e,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
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
    <form action="/loginprocess" method="POST">
    <input type="text" name="em" placeholder="EMAIL">
    <input type="text" name="pw" placeholder="PASSWORD">
    <input type="submit">    
</form>
</body>
</html>
	`)
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
	}


	u, ok := db[e]
	if !ok {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	if p != u.Password {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	c := &http.Cookie{
		Name:  "user-session",
		Value: e,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{
			Name:  "user-session",
			Value: "JamesBondUUID",
			Path:  "/",
		}
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
