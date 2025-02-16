package main

import "fmt"

var globalBlock = 125

func main() {

	fmt.Println(globalBlock)
	{
		localBlock := 124
		fmt.Println(localBlock)
		{
			localBlock := 1250
			fmt.Println(localBlock)
		}
		fmt.Println(localBlock)
	}
}
