package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing tempalte: %w", err)
	}

	return Template{
		HTMLTpl: tpl,
	}, nil
}

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.HTMLTpl.Execute(writer, nil)

	if err != nil {

		log.Printf("executing template: %v", err)

		http.Error(writer, "There was an error executing the template.", http.StatusInternalServerError)

		return
	}
}

func (t Template) Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		HTMLTpl: tpl,
	}, nil
}
