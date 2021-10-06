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
		people := []Person{
			{"Bob", "Smith"},
			{"Larry", "Flint"},
		}

		tp.Execute(rw, people)
	})

	http.HandleFunc("/longlist/", func(rw http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)
		people := []Person{
			{"Bob", "Smith"},
			{"Larry", "Flint"},
			{"Susan", "Sarandon"},
			{"Brad", "Pitt"},
			{"Betty", "White"},
		}

		tp.Execute(rw, people)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles()
		tp.Execute(rw, nil)
	})

	http.ListenAndServe(":8001", nil)
}
