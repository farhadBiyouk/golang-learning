package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name      string `json:"first_name" xml:"name"`
	Age       int    `json:"age" xml:"age"`
	Family    string `json:"last_name" xml:"familyyyyyyyy"`
	Gender    string `json:"gender" xml:"sex"`
	Height    int    `json:"height" xml:"ghad"`
	Weight    int    `json:"weight" xml:"vazen"`
	HairColor string `json:"hair_color,omitempty" xml:"color_mo,omitempty"`
}

func main() {
	person := Person{
		Name:   "ali",
		Age:    25,
		Family: "hassani",
		Gender: "male",
		Height: 186,
		Weight: 110,
		// HairColor: "Black",
	}
	personJson, _ := json.MarshalIndent(person, "", "	")
	personXml, _ := xml.MarshalIndent(person, "", "	")
	fmt.Println(string(personJson))
	fmt.Println(string(personXml))

}
