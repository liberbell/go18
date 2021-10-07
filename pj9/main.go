package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	templateFile := "personlist.tmpl"

	http.HandleFunc("/shortlist/", func(rw http.ResponseWriter, r *http.Request) {
		tp, _ := template.ParseFiles(templateFile)
	})
}
