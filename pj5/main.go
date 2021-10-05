package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	tmp1 := "<B><font color='green'>First Name: </font></B> {{.FirstName}} <BR> <B><font color='red'>Last Name: </font></B> {{.LastName}}"

	router := httprouter.New()
	router.GET("/name/:firstName/:lastName", func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		t := template.New("NameTemplate")
		tp, _ := t.Parse(tmp1)

		tp.Execute(rw, &Person{ps.ByName("firstName"), ps.ByName("lastname")})
	})

	http.ListenAndServe(":8001", nil)
}
