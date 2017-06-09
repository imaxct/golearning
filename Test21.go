package main

import (
	"net/http"
	"log"
	"io"
)

func testHandler(res http.ResponseWriter, req *http.Request) {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	io.Copy(res, r.Body)
}

func main() {

	http.HandleFunc("/", testHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
