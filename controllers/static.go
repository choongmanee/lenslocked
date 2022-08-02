package controllers

import (
	"html/template"
	"net/http"

	"github.com/choongmanee/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl.Execute(writer, nil)
	}
}

func FAQHandler(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{Question: "Who?", Answer: "Chung"},
		{Question: "What?", Answer: "Man"},
		{Question: "When?", Answer: "Now"},
		{Question: "Where?", Answer: "Here"},
		{Question: "Why?", Answer: "Because"},
		{Question: "How?", Answer: "Chung"},
		{Question: "What are your support hours?", Answer: "We have support staff answering emails 24/7, though response times may be a bit slower on weekends."},
		{Question: "Is there a free version?", Answer: "Yes! we offer a free trial for 30 days on any paid plans"},
		{Question: "How do I contact support?", Answer: " Email use - <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>."},
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		tpl.Execute(writer, questions)
	}
}
