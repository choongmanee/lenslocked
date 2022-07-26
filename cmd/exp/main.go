package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
	}{
		Name: "John Smith",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
