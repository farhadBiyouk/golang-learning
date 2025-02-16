package main

import "fmt"

func main() {

	i, j := 10, 20

	var ip *int = &i
	var jp *int = &j

	fmt.Println("Address of i:", &i)
	fmt.Println("Address of ip  pointer:", ip)

	fmt.Println("Address of j:", &j)
	fmt.Println("Address of jp pointer:", jp)

	// value type
	i2 := i
	i2 = i2 + 2

	fmt.Println("value of i after increase i2", i)
	fmt.Println("value of i2 after increase i2", i2)

	// reference type
	var ip2 *int = &j
	*ip2 = *ip2 + 3
	fmt.Println("value ip2 is: ", *ip2)
	fmt.Println("value i is: ", j)
}
