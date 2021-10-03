package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			r.ParseForm()
			fmt.Printf("First Name from Form: %s\n", r.Form("firstname"))
			fmt.Printf("Last Name from Form: %s\n", r.Form("lastname"))

			fmt.Printf("First Name from PostForm: %s\n", r.PostForm("firstname"))
			fmt.Printf("Last Name from PostForm: %s\n", r.PostForm("lastname"))
		}
	})
}
