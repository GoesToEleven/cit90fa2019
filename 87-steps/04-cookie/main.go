package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Value string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// set cookie
	c := &http.Cookie{
		Name: "my-cook",
		Value: "some value here",
		Path: "/",
	}
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	// retrieve cookie
	c, err := r.Cookie("my-cook")
	if err != nil {
		http.Error(w, "couldn't get my-cook cookie", http.StatusBadRequest)
		return
	}

	pd := pageData{
		Value: c.Value,
	}
	

	// output cookie values
	tpl.ExecuteTemplate(w, "bar.gohtml", pd)
}
