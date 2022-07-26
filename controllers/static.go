package controllers

import (
	"net/http"

	"github.com/choongmanee/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl.Execute(writer, nil)
	}
}
