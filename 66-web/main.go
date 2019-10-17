package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var headers string
	for k, xs := range r.Header {
		for _, s := range xs {
			headers += k + s
			headers += "\n"
		}
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("couldn't ioutil.Readall", err)
	}

	fmt.Fprintf(w, "METHOD: %s\nURL: %s\nPROTO:%s\nHEADERS: %s\nBODY: %s\nCONTENT LENGTH: %d\nHOST: %s\nREQUEST URI: %s\nUSER AGENT: %s\nREFERRER: %s\n", r.Method, r.URL.String(), r.Proto, headers, string(bs), r.ContentLength, r.Host, r.RequestURI, r.UserAgent(), r.Referer())
}
