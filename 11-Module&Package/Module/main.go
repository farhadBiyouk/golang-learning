package main

import (
	"fmt"

	jajlli "github.com/jalaali/go-jalaali"
)

func main() {
	y, m, d, err := jajlli.ToGregorian(1403, 15, 04)
	if err == nil {
		fmt.Println(y, m, d)
	} else {
		fmt.Println(err)
	}

	// article1 := blogApi.Blog{ID: 1, Title: "first title", Author: "farhad", Description: "this is a first my blog"}
	// fmt.Println(article1)
}
