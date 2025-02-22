package main

import "fmt"

type Person struct {
	ID     int
	Name   string
	Family string
}

type PersonOption func(*Person)

func main() {
	person := NewPerson(SetId(1), SetName("hassan"), SetFamily("abbasi"))
	fmt.Println(person)
}

func NewPerson(options ...PersonOption) *Person {
	ps1 := &Person{ID: 1}
	for _, opt := range options {
		opt(ps1)
	}
	return ps1
}
func SetId(id int) PersonOption {
	return func(person *Person) {
		person.ID = id
	}
}

func SetName(name string) PersonOption {
	return func(person *Person) {
		person.Name = name
	}
}

func SetFamily(family string) PersonOption {
	return func(person *Person) {
		person.Family = family
	}
}
