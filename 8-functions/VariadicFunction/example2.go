package main

import "fmt"

func main() {
	Logs("ali", 1, 2, 3, true)
}

func Logs(log ...interface{}) {
	for _, item := range log {
		fmt.Printf("%v \n", item)
	}
}
