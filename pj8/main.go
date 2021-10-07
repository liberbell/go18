package main

import "flag"

func main() {
	var port, message string

	flag.StringVar(&port, "port", "8001", "The HTTP port")
	flag.StringVar(&message, "message", "", "The response message")

	flag.Parse()

}
