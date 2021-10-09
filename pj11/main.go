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
			{"Bo b", "Sm i  th"},
			{"La   rry", "F   lint"},
			{"Su san", "Sa ra  ndon"},
			{"Bra   d", "P  itt"},
			{"Bet  ty", "W  hi te"},
		}

		tp.Execute(rw, people)
	})
	http.ListenAndServe(":8001", nil)
}
