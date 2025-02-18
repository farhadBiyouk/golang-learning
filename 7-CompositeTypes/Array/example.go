package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3}
	ChangeValue(&numbers)
	fmt.Println(numbers)
}
func ChangeValue(lst *[3]int) {
	lst[1] = 10
	fmt.Println(*lst)
}
