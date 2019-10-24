package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title   string
	Heading string
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "HOME",
		Heading: "Welcome to our site as if it's 1999",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "ABOUT",
		Heading: "ABOUT US",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}

func shop(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "SHOP",
		Heading: "BUY SOME STUFF",
	}

	err := tpl.ExecuteTemplate(w, "shop.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}

}
func contact(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "CONTACT",
		Heading: "What? What?",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}
