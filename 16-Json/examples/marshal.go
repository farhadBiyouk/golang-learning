package examples

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Age    int    `json:"age,omitempty"`
}

func ConvertToJsonWithMarshal() {

	p1 := &Person{Name: "Ali", Family: "Rezaee", Age: 20}
	p2 := &Person{Name: "Iman", Family: "Mohammadi", Age: 25}

	p1Json, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	p2Json, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(p1Json))
	fmt.Println(string(p2Json))
}

func ConvertToJsonWithMarshal2() {

	p1 := &Person{Name: "Ali", Family: "Rezaee", Age: 20}
	p2 := &Person{Name: "Iman", Family: "Mohammadi", Age: 25}

	persons := []Person{*p1, *p2}

	PersonJson, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(PersonJson))
}
