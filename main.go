package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func showHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	fmt.Fprint(writer, id)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Route("/middleware", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "<h1>middleware</h1>")
		})
	})
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/show/{id}", showHandler)

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", r)
}
