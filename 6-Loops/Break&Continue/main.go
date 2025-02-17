package main

import "fmt"

func main() {
	// for i := 1; i < 100; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	for i := 1; i < 100; i++ {
		if i == 50 {
			fmt.Println("Visit number 50")
			break
		}

	}

}
