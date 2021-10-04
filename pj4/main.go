package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

type PersonJsonField struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type PersonJsonMarshaler struct {
	FirstName string
	LastName  string
}

func (p *PersonJsonMarshaler) MarshalJSON() ([]byte, error) {
	return []byte("{\"Name\":\"" + p.FirstName + " " + p.LastName + "\"}"), nil
}

func main() {
	http.HandleFunc("/json/default/", func(rw http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&Person{FirstName: "Bob", LastName: "Smith"})
		rw.Write(j)
	})

	http.HandleFunc("/json/fields/", func(rw http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonField{FirstName: "Bob", LastName: "Smith"})
		rw.Write(j)
	})

	http.HandleFunc("/json/marshal/", func(rw http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonMarshaler{FirstName: "Bob", LastName: "Smith"})
		rw.Write(j)
	})

	http.ListenAndServe(":8001", nil)
}
