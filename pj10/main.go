package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func main() {
	templateFile := "personlist.tpl"

	fmap := template.FuncMap{
		"removeSpace": removeSpaces,
		"capitalize":  strings.ToUpper,
	}

	http.HandleFunc("/shortlist", func(rw http.ResponseWriter, r *http.Request) {
		tp, _ := template.New("personlist.tpl").Funcs(fmap).ParseFiles(templateFile)
	})
}
