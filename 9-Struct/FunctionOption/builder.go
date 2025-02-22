package main

import "fmt"

type Person struct {
	ID     int
	Name   string
	Family string
}

type PersonBuilder struct {
	Person
}

func main() {

	personBuilder := &PersonBuilder{}
	result := personBuilder.SetId(1).SetName("Matyam").SetFamily("Khan").Build()
	fmt.Println(result)

}

func (builder *PersonBuilder) SetId(id int) *PersonBuilder {
	builder.ID = id
	return builder
}

func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}

func (builder *PersonBuilder) SetFamily(family string) *PersonBuilder {
	builder.Family = family
	return builder
}

func (builder *PersonBuilder) Build() Person {
	return builder.Person
}
