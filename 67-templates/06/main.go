package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	xs := []string{"buddha", "gandhi", "mlk", "santa"}

	err := tpl.ExecuteTemplate(w, "index.gohtml", xs)
	if err != nil {
		http.Error(w, "couldn't run template", http.StatusInternalServerError)
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	m := map[int]string{}
	m[1] = "jenny"
	m[2] = "james"

	err := tpl.ExecuteTemplate(w, "about.gohtml", m)
	if err != nil {
		http.Error(w, "couldn't run template", http.StatusInternalServerError)
	}
}
