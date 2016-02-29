package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s.Who+s.Punct+s.Greeting)
}

func main() {
	http.Handle("/string", String("Hello, Welcome"))
	http.Handle("/struct", &Struct{"Hello", ":", "Sam"})
	http.ListenAndServe(":4000", nil)
}
