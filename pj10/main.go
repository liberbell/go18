package main

import "strings"

type Person struct {
	FirstName string
	LastName  string
}

func removeSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func main() {
	templateFile := "personlist.tpl"
}
