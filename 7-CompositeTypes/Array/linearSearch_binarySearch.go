package main

import "fmt"

func main() {

	numbers := [8]int{10, 14, 16, 9, 36, 75, 44, 99}

	// fmt.Println(binarySearch(&numbers, 375, 0, len(numbers)-1))
	linearSearch(&numbers)
}

func binarySearch(lst *[8]int, target int, low int, high int) int {

	if low > high {
		return 0
	} else {
		mid := (low + high) / 2
		if lst[mid] == target {
			return 1
		} else if target < lst[mid] {
			return binarySearch(lst, target, low, mid-1)
		} else {
			return binarySearch(lst, target, mid+1, high)
		}
	}
}

func linearSearch(lst *[8]int) {
	var num_for_search int
	print("Enter number for search? ")
	fmt.Scanln(&num_for_search)

	for index, item := range lst {
		if item == num_for_search {
			println(item, "founded in index", index)
			break
		}
	}
}
