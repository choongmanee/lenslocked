package main

import (
	"fmt"
	"net/http"

	"github.com/choongmanee/lenslocked/controllers"
	"github.com/choongmanee/lenslocked/templates"
	"github.com/choongmanee/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	homeTpl := views.Must(views.ParseFS(templates.FS,
		"layout-page.gohtml",
		"partials.gohtml",
		"home.gohtml",
	))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.ParseFS(templates.FS,
		"layout-page.gohtml",
		"partials.gohtml",
		"contact.gohtml",
	))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.ParseFS(templates.FS,
		"layout-page.gohtml",
		"partials.gohtml",
		"faq.gohtml",
	))
	r.Get("/faq", controllers.FAQHandler(faqTpl))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS,
		"layout-page.gohtml",
		"partials.gohtml",
		"signup.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "page not found", http.StatusNotFound)
	})

	fmt.Println("starting server on :3000")
	http.ListenAndServe(":3000", r)
}
