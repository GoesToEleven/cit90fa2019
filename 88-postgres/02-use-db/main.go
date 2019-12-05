package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	// setup database
	var err error
	connStr := "user=bond password=password dbname=wishlists sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	// rest of code
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	type person struct {
		First string
		Last  string
	}

	people := []person{}

	rows, err := db.Query("SELECT pfirstname, plastname from people;")
	if err != nil {
		log.Panicln("db.query blew up", err.Error())
		http.Error(w, "db query blew up"+err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var p person
		err = rows.Scan(&p.First, &p.Last)
		if err != nil {
			log.Panicln("rows.Scan blew up", err.Error())
			http.Error(w, "rows scan blew up"+err.Error(), http.StatusInternalServerError)
			return
		}
		people = append(people, p)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", people)
}
