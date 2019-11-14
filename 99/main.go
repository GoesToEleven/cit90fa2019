package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

// User is exported so that it can be an embedded type in pageData
type User struct {
	First    string
	Email    string
	Password []byte
	LoggedIn bool
}

type session struct {
	email   string
	expires time.Time
}

type pageData struct {
	User
	Title   string
	Heading string
}

// create variable to hold value of TYPE map[string]user
// DOES NOT INITIALIZE variable with VALUE in memory
// var db map[string]user

// create variable to hold value of TYPE map[string]user
// INITIALIZES variable with VALUE in memory
// key is email
// value is user
var db = map[string]User{}

// sesh
// key is unique signed sessionID
// value is email
// why use a session id?
// why not just use a userID such as email or some other userID?
// answer: we want to be able to invalidate a sessionID
var sesh = map[string]session{}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
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

	u := User{}
	ok := true
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "Acme Inc.",
		Heading: "Welcome To Acme Inc.",
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("index.gohtml couldn't ExecuteTemplate", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	ok := false
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "About Acme Inc.",
		Heading: "All About Acme",
	}

	err = tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println("about.gohtml couldn't ExecuteTemplate", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	ok := false
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "Contact Acme Inc.",
		Heading: "Contact Acme Incorporated",
	}

	err = tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println("contact.gohtml couldn't ExecuteTemplate", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	if _, ok := db[c.Value]; ok {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "SIGN UP",
		Heading: "SIGN UP AT ACME INC.",
	}

	err = tpl.ExecuteTemplate(w, "signup.gohtml", pd)
	if err != nil {
		log.Println("signup.gohtml couldn't ExecuteTemplate", err)
	}
}

func processSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	f := r.FormValue("fn")
	if f == "" {
		http.Error(w, "first name cannot be empty", http.StatusBadRequest)
		return
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
		return
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
		return
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "password could not be bcrypt'd", http.StatusBadRequest)
		return
	}

	// never directly store a password
	// store the bcrypt or scrypt of the password
	u := User{
		First:    f,
		Email:    e,
		Password: bs,
	}

	db[e] = u

	// now we want to store some session ID
	// this session ID should be signed
	// signing it means we can check to see if it was changed on the user side
	// if it was changed on the user side, we won't accept it
	// for session id, how about we combine their email with a secret key then hash that
	// so that will be some unique string
	// email|<hash of email+secretKey>

	c := &http.Cookie{
		Name:  "user-session",
		Value: e,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	if _, ok := db[c.Value]; ok {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "LOGIN",
		Heading: "LOGIN TO ACME INC.",
	}

	tpl.ExecuteTemplate(w, "login.gohtml", pd)
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

	u := User{}
	ok := false
	if u, ok = db[e]; !ok {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		http.Error(w, "username or password not correct BCRYPT'D", http.StatusBadRequest)
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
			Name: "user-session",
		}
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
