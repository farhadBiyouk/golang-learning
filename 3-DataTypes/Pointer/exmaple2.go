package main

import "fmt"

func main() {

	number := 10
	fmt.Println("address var in main block is: ", &number)
	ChangeNumberByRef(&number)

	fmt.Println("value is:", number)

	ChangeNumberByValue(number)
	fmt.Println("value is: ", number)

}

func ChangeNumberByRef(num *int) {
	println("changeNumberByRef Address of num:", num)
	*num = *num + 2
}

func ChangeNumberByValue(num int) {
	println("changeNumberByValue Address of num:", &num)
	num = num + 3
}
