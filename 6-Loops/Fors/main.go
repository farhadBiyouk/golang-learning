package main

import "fmt"

func main() {

	// Infinite loop
	// for {
	// 	fmt.Println("Hello golang")
	// }

	// like while
	// i := 0
	// for i <= 10{
	// 	fmt.Println("Hello Go")
	// 	i +=1 // i++
	// }

	// for j:=0; j<11;j++{
	// 	fmt.Println("go is very good")
	// }

	// usage for slice and array

	// numbers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// for index, item := range numbers {
	// 	fmt.Printf("value %d in index %d.\n", item, index)
	// }

	//nested loops
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "    ")
		}
		fmt.Println(" ")
	}
}
