package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	// index out of range
	// numbers := []int{1, 2}
	// showSlice(numbers, 0)
	// panic
	// panic("panic")

	// log.Fatal("fatal")

	defer func() {
		if err := recover(); err != nil {
			fmt.Sprintf("have error %s", err)
			// debug.PrintStack()
			debug.Stack()
		}
	}()

	Divide(10, 0)
}

// func showSlice(number []int, index int) {
// 	fmt.Println(number[index])
// }

func Divide(a, b int) {
	defer func() {
		fmt.Println("Defer of divide")
	}()
	fmt.Printf("Start of divide: %d , %d\n", a, b)
	fmt.Printf("Result: %d\n", a/b)
	fmt.Printf("End of divide: %d , %d\n", a, b)

}
