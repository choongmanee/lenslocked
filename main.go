package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello world foo</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
