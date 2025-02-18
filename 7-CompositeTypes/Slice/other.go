package main

import "fmt"

// append new data in special index
// remove data from  special index

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	numbers = append(numbers, 0)
	copy(numbers[8:], numbers[7:])
	numbers[7] = 85
	fmt.Println(numbers)

	numbers = append(numbers[:7], numbers[8:]...)
	fmt.Println(numbers)
	numbers = numbers[:0]
}
