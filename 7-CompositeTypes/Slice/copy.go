package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 5}
	number2 := []int{10, 20, 30, 40, 50}

	count := copy(number1, number2)
	fmt.Println(count)
	fmt.Println(number1)
}
