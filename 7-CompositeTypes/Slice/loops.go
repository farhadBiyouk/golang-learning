package main

import (
	"fmt"
	"strings"
)

func main() {

	names := []string{"farhad", "ali", "reza", "asma"}

	for index, item := range names {
		names[index] = strings.ToUpper(item)

	}
	fmt.Println(names)
}
