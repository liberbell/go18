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
		people := []Person{Person{"Bob", "Smith"}, Person{"Larry", "Flint"}}
		tp.Execute(rw, people)
	})

	http.HandleFunc("/longlist/", func(rw http.ResponseWriter, r *http.Request) {
		tp, _ := template.ParseFiles(templateFile)
		people := []Person{Person{"Bob", "Smith"}, Person{"Larry", "Flint"}, Person{"Susan", "Sarandon"}, Person{"Brad", "Pitt"}, Person{"Betty", "White"}}
		tp.Execute(rw, people)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tp, _ := template.ParseFiles(templateFile)
		tp.Execute(rw, nil)
	})

	http.ListenAndServe(":8001", nil)
}
