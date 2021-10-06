package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	args := os.Args
	port := os.Args[1]
	message := os.Args[2]

	fmt.Println(args)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(message))
	})

	http.ListenAndServe(":"+port, nil)
}
