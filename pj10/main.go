package main

import (
	"html/template"
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
}
