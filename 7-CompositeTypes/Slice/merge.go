package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 5}
	number2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	numbers := []int{}

	fmt.Println(number1)
	numbers = append(number1, number2...)
	fmt.Println(numbers)

}
