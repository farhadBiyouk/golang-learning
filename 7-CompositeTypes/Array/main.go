package main

import "fmt"

func main() {

	//  define array

	// var numbers = [3]int{1, 2, 3}
	// var number1 [4]int
	// var number2 = [...]string{}
	var new_name string
	names := [6]string{}
	for i := 0; i < len(names); i++ {
		print("Enter name: ")
		fmt.Scanln(&new_name)
		names[i] = new_name
	}
	for _, item := range names {
		fmt.Println(item)
	}
}
