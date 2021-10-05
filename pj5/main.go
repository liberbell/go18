package main

import (
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	tmp1 := "<B><font color='green'>First Name: </font></B> {{.FirstName}} <BR> <B><font color='red'>Last Name: </font></B>"

	router := httprouter.New()
	router.GET("/name/:firstname/:lastname", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params ps httprouhttprouter.Param) {
		t := template.New("NameTemplate")
		tp, _ := t.Parse(tmp1)
	})
}
