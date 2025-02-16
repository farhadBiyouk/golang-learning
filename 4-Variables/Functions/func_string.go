package main

import (
	"fmt"
	"strings"
)

func main() {
	str_test := "this is  golang tutorial go go"

	fmt.Println(strings.Contains(str_test, "go"))
	fmt.Println(strings.ContainsAny(str_test, "1"))

	fmt.Println(strings.Count(str_test, "go"))

	fmt.Println(strings.Cut(str_test, "go"))

	result := strings.Split(str_test, " ")
	println("word count: ", len(result))

	for _, item := range result {
		println(item)
	}

	result2 := strings.Join(result, "-")
	println(result2)

	println(strings.Repeat("farhad ", 5))

	println(strings.Replace(str_test, "go", "golang", 1))
	println(strings.ReplaceAll(str_test, "go", "golang"))

	println(strings.Compare("golang", "golang"))
	println(strings.EqualFold("golang", "Golang"))

	println(strings.HasPrefix("Iran", "Ir"))
	println(strings.HasSuffix("Iran", "n"))

	println(strings.Index("Iran", "n"))

	println(strings.ToLower("IRAN"))
	println(strings.ToUpper("Iran"))
	println(strings.Title("Iran is a good country"))
	println(strings.ToTitle("Iran is a good country"))

	println(strings.Trim("         Iran is a good country        ", " "))
	println(strings.TrimRight("Iran is a good country !!!!", " !"))
	println(strings.TrimLeft("?? Iran is a good country", "?? "))

}
