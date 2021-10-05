package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	templateFile := "personlist.html"

	http.HandleFunc("/shortlist/", func(rw http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)
		people := []Person{
			{"Bob", "Smith"},
			{"Larry", "Flint"},
		}

		tp.Execute(rw, people)
	})
}
