package main

import "fmt"

func main() {

	// define slice
	var names = []string{}
	families := []string{}
	ages := make([]int, 8) // 8 is length
	jobs := make([]string, 8, 16)

	// create slice with array

	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	num2 := numbers[0:6]

	numbers[2] = 20

	// fmt.Printf("%v\n", numbers)
	// fmt.Printf("%v\n", num2)

	// fmt.Println("num2 length: ", len(num2))
	// fmt.Println("num2 cap: ", cap(num2))

	// fmt.Println("numbers length: ", len(numbers))
	// fmt.Println("numbers cap: ", cap(numbers))

	num2 = append(num2, 99)
	num2 = append(num2, 98)
	num2 = append(num2, 97)
	num2 = append(num2, 96, 95, 94, 93, 92, 91, 90)

	numbers[7] = 101

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", num2)

	fmt.Println("num2 length: ", len(num2))
	fmt.Println("num2 cap: ", cap(num2))

	fmt.Println("numbers length: ", len(numbers))
	fmt.Println("numbers cap: ", cap(numbers))
}
