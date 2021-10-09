package main

import (
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	templateFile := "personlist.tpl"

	http.HandleFunc("/longlist/", func(rw http.ResponseWriter, r *http.Request) {

	})
}
