package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/plaintext/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString()
	})
}
