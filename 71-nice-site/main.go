package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	Title string
	Heading string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	d := pageData{
		Title: "ACME CO",
		Heading: "It's Your Party",
	}
	tpl.ExecuteTemplate(w, "index.gohtml", d)
}
func about(w http.ResponseWriter, r *http.Request)   {
		d := pageData{
		Title: "ABOUT",
		Heading: "ABOUT OUR COMPANY",
	}
	tpl.ExecuteTemplate(w, "about.gohtml", d)
}

func shop(w http.ResponseWriter, r *http.Request)    {
		d := pageData{
		Title: "SHOP",
		Heading: "Feel Good. Buy Stuff.",
	}
	tpl.ExecuteTemplate(w, "shop.gohtml", d)
}

func contact(w http.ResponseWriter, r *http.Request) {
		d := pageData{
		Title: "CONTACT",
		Heading: "Don't Actually Bother Us",
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", d)
}

