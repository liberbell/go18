package main

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

func main() {
	a
}
