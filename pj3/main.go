package main

import "net/http"

func main() {
	http.HandleFunc("/notfound/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(404)
	})

	http.HandleFunc("/servererror/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(500)
	})

	http.HandleFunc("/customheader/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("custom-header", "This is my custom header.")
	})
}
