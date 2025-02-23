package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Sleep()
	Walk()
}

type Human interface {
	Animal
	Speak()
	Think()
}

type Employee struct {
	Human
	Name string
	Age  int
}

type Dog struct {
	Name string
}

func main() {
	emp1 := &Employee{Name: "Fariborz", Age: 46}
	dog1 := Dog{Name: "sarabi"}

	var human Human = emp1
	var animal Animal = &dog1

	human.Eat()
	human.Sleep()
	human.Walk()
	human.Speak()
	human.Think()
	fmt.Println("_____________________")

	animal.Eat()
	animal.Sleep()
	animal.Walk()
}

func (emp *Employee) Eat() {
	fmt.Println("Human Eating")
}
func (emp *Employee) Sleep() {
	fmt.Println("Human Sleeping")
}
func (emp *Employee) Walk() {
	fmt.Println("Human Walking")
}
func (emp *Employee) Speak() {
	fmt.Println("Human Speaking")
}
func (emp *Employee) Think() {
	fmt.Println("Human Thinking")
}

func (emp *Dog) Eat() {
	fmt.Println("Dog Eating")
}

func (emp *Dog) Sleep() {
	fmt.Println("Dog Sleeping")
}
func (emp *Dog) Walk() {
	fmt.Println("Dog Walking")
}
