package main

import "fmt"

func main() {

	say_hello()
	get_name("ali")
	result := add()
	print(result, "\n")
	fmt.Println(div(20, 2))
	fmt.Println(mul(5, 5))
}

// define function without input and output
func say_hello() {
	fmt.Println("Hello world!!!")
}

// define function with input and without output
func get_name(name string) {
	fmt.Println("hello", name)
}

// define function without input and with output
func add() int {
	return 12 + 13
}

// define function with input and output
func div(num1 int, num2 int) int {
	return num1 / num2
}

// define function with input and multiple output
func mul(num1 int, num2 int) (numb1 int, numb2 int, result int) {
	numb1 = num1
	numb2 = num2
	result = num1 * num2
	return
}
