package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.mongard.iir")
	if err != nil {
		fmt.Println("an error has occurred from get link ")
	}
	fmt.Println(response)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("an error has occurred from response body ")
	}
	fmt.Println(string(responseBody))
}
