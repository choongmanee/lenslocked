package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/choongmanee/lenslocked/controllers"
	"github.com/choongmanee/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	homeTpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", r)
}
