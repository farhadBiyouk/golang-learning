package main

import "fmt"

type Number interface {
	int | float64
}

func main() {

	number1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sum(number1))

	numbers2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	fmt.Println(Sum(numbers2))

}

func Sum[T Number](slc []T) (sum T) {
	for _, item := range slc {
		sum += item
	}
	return
}
