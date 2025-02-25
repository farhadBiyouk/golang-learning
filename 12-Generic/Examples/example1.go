package main

import "fmt"

type NumberTYPE interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func main() {

	fmt.Println(SumTowNumbers(1, 2))
	fmt.Println(SumTowNumbers2(1.1, 2.2))
	fmt.Println("__________________________")
	fmt.Println(SumTwoNumbersWithAnyType(1, 2))
	fmt.Println(SumTwoNumbersWithAnyType(1.1, 2.2))
	fmt.Println(SumTwoNumbersWithAnyType("Hello", "World"))
}

func SumTowNumbers(a int, b int) int {
	return a + b
}
func SumTowNumbers2(a float64, b float64) float64 {
	return a + b
}

func SumTwoNumbersWithAnyType[T NumberTYPE](a T, b T) T {
	return a + b
}
