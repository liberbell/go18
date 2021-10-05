package main

import "github.com/julienschmidt/httprouter"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	tmp1 := "<B><font color='green'>First Name: </font></B> {{.FirstName}} <BR> <B><font color='red'>Last Name: </font></B>"

	router := httprouter.New()
}
