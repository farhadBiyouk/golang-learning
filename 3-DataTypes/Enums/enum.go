package main

import (
	"encoding/json"
	"fmt"
)

// define constant
const PI = 3.14

// define enum

type Order struct {
	Id     int
	Name   string
	Price  int
	Status OrderType
}

type OrderType int

const (
	Created        OrderType = 0
	Processing     OrderType = 1
	WaitForPayment OrderType = 2
	PaymentDone    OrderType = 3
	Issue          OrderType = 4
	Refund         OrderType = 5
)

func main() {

	order1 := Order{Id: 1, Name: "laptop", Price: 1000, Status: PaymentDone}
	order2 := Order{Id: 2, Name: "mobile", Price: 993, Status: Created}

	order1json, _ := json.Marshal(order1)
	order2json, _ := json.Marshal((order2))

	fmt.Println(string(order1json))
	fmt.Println(string(order2json))
}
