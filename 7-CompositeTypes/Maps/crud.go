package main

import "fmt"

type Person struct {
	Name   string
	Family string
	Age    int
}

func main() {
	names := make(map[int]Person)

	// create
	names[1234567890] = Person{Name: "Farhad", Family: "Biyouk", Age: 25}
	names[1234567891] = Person{Name: "Ramin", Family: "Jamshdi", Age: 27}
	names[1234567892] = Person{Name: "Ali", Family: "Amini", Age: 23}

	// retrieve
	fmt.Println(names[1234567890])

	// update
	names[1234567892] = Person{Name: "Amin", Family: "alavai", Age: 23}
	fmt.Println(names[1234567892])

	//delete
	delete(names, 1234567891)
	fmt.Println(names)

}
