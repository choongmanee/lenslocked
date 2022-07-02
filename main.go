package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>home</h1>")
}

func contactHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>contact</h1>")
}

func faqHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>faq</h1>")
}
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", r)
}
