package main

import (
	"errors"
	"fmt"
)

func main() {

	num, err := CreateErrorMethod2(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

func CreateErrorMethod1(number int) (int, error) {
	if number == 0 || number < 0 {
		return 0, errors.New("number is invalid from method 1")
	}
	return number * 5, nil
}

func CreateErrorMethod2(number int) (int, error) {
	if number == 0 || number < 0 {
		return 0, fmt.Errorf("number is invalid: %d", number)
	}
	return number * 6, nil
}
