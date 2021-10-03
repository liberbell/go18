package main

import (
	"fmt"
	"io/ioutil"
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

		http.ServeFile(w, r, "PersonForm.html")
	})

	http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			r.ParseMultipartForm(0)
			fmt.Printf("File information: %s\n", r.MultipartForm)

			file, _, _ := r.FormFile("file")
			bytes, _, := ioutil.ReadAll(file)

			fmt.Printf("\n\nFile contents: %s\n", bytes)
		}
		http.ServeFile()
	})
}
