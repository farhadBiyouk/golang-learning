package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type PersonOption struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// create instance from struct

	// method 1
	person1 := Person{FirstName: "Ali", LastName: "Khan", Age: 25}

	// method 2
	person2 := new(Person)
	person2.FirstName = "Maryam"
	person2.LastName = "Khan"
	person2.Age = 33

	//method 3
	person3 := &Person{FirstName: "Amir", LastName: "Khan", Age: 25}

	// method 4
	person4 := NewPerson("Reza", "Khan", 23)

	// method5

	person5 := NewPersonWithOption(PersonOption{FirstName: "Hemed", LastName: "Khan", Age: 35})

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
	fmt.Println(person5)
}

// method 4  ==> like constructor
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}

// method 4  ==> like constructor
func NewPersonWithOption(person PersonOption) *Person {
	return &Person{FirstName: person.FirstName, LastName: person.LastName, Age: person.Age}
}
