package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello from inside function first")
	}()

	insid := func() {
		fmt.Println("Hello from inside function second")
	}
	insid()

	fmt.Println(func(names ...int) int {
		total := 0
		for _, item := range names {
			total += item
		}
		return total
	}(1, 1, 5, 17, 85, 86, 326, 78, 51, 61, 61, 6, 14, 855, 6, 6))
}
