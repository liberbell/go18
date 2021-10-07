package main

import (
	"flag"
	"net/http"
)

func main() {
	var port, message string

	flag.StringVar(&port, "port", "8001", "The HTTP port")
	flag.StringVar(&message, "message", "", "The response message")

	flag.Parse()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(message))
	})

	http.ListenAndServe(":"+port, nil)
}
