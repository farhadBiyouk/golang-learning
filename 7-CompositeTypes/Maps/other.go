package main

import "fmt"

type Person struct {
	Name   string
	Family string
	Age    int
}

func main() {

	info := make(map[int]Person)

	sort_info := []Person{}

	info[1234567891] = Person{Name: "Ramin", Family: "Jamshdi", Age: 27}
	sort_info = append(sort_info, info[1234567891])
	info[1234567892] = Person{Name: "Ali", Family: "Amini", Age: 23}
	sort_info = append(sort_info, info[1234567892])
	info[1234567890] = Person{Name: "Farhad", Family: "Biyouk", Age: 25}
	sort_info = append(sort_info, info[1234567890])

	for _, item := range sort_info {
		fmt.Println(item)
	}
}
