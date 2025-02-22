package main

import "fmt"

// define anonymous
type Person struct {
	int
	string
	float64
}

func main() {
	person := &Person{1, "Reza abbasi", 26}
	fmt.Println(*person)

	// Api call and get response and define anonymous struct
	api_response := struct {
		StatusCode        int
		MessageStatusCode string
		TransactionAmount float64
		TransactionTime   string
	}{
		StatusCode:        200,
		MessageStatusCode: "Ok",
		TransactionAmount: 100,
		TransactionTime:   " 2025-05-23T12:35:45",
	}

	fmt.Println(api_response)
}
