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
	templateFile := "personlist.tpl"

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
	http.ListenAndServe(":8001", nil)
}
