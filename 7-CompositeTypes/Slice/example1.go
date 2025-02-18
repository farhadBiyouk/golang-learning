package main

import "fmt"

func main() {
	// send slice to func
	// send slice to loop and problem
	// copy slice
	// merge two slice
	numbers := []int{1, 2, 3, 4, 5, 6}

	AddItem(&numbers)
	ChangeValue(numbers)

	fmt.Println(numbers)
}

func ChangeValue(lst []int) {
	for i := 0; i < len(lst); i++ {
		lst[i] *= 3
	}
}

func AddItem(lst *[]int) {
	*lst = append(*lst, 7)

}
