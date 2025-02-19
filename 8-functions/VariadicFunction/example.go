package main

import "fmt"

func main() {

	fmt.Println(calculator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func calculator(number ...int) (sum int, mul int) {
	mul = 1

	for _, item := range number {
		mul *= item
		sum += item
	}
	return
}
