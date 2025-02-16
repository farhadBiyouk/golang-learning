package main

import "fmt"

func main() {

	// define with var

	var number1 int = 10
	var name string = "ali"
	var age int
	var is_married bool
	var pi = 3.14

	//define with walrus

	family := "parsa"
	mobile := 8746511239

	// define multiple

	var name, family, age = "ali", "alavi", 32
	name, family, age := "ali", "alavi", 32
	var (
		name   = "ali"
		family = "alavi"
		age    = 32
	)

	// fmt.Println(name, family, age)

}
