package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("Im fine"))
	http.Handle("/struct", &Struct{"Hello", ":", "World"})
	err := http.ListenAndServe("localhost:4440", nil)
	if err != nil {
		log.Fatal(err)
	}
}
