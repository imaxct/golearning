package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	res := make(chan string)
	go func(url string, c chan string) {
		r, err := http.Get(url)
		if err != nil {
			res <- err.Error()
			return
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		c <- string(body)
	}("https://golang.org/",res)
	fmt.Println(<-res)
}