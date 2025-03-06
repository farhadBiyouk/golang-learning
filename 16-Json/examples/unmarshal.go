package examples

import (
	"encoding/json"
	"fmt"
)

func FetchDataFromJson() {
	person := []byte(`{"name":"Ali","family":"Rezaee","age":20}`)

	var personType = Person{}
	err := json.Unmarshal(person, &personType)
	if err != nil {
		panic(err)
	}

	fmt.Println(personType)
}

func FetchDataFromJson2() {
	person := []byte(`[{"name":"Ali","family":"Rezaee","age":20},{"name":"Milad","family":"Hassani","age":21}]`)

	var personType = []Person{}
	err := json.Unmarshal(person, &personType)
	if err != nil {
		panic(err)
	}

	fmt.Println(personType)
}
