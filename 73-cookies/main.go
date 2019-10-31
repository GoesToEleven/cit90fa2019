package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{
			Name:  "user-session",
			Value: "JamesBondUUID",
			Path:  "/",
		}
	}

	http.SetCookie(w, c)
	
	w.Header().Set("content-type","text/html; charset=utf-8")
	io.WriteString(w, "<h1>COOKIE:"+c.Name+" "+c.Value+"</h1>")
}
