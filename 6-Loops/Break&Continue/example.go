package main

import (
	"fmt"
	"math/rand"
)

type CreditCard struct {
	CardNumber string
	ExpireTime string
	Cvv2       int
	BankName   string
}

func getRemainAmountAccount(cardNum string, expire string) int {
	min := 100_000
	max := 3_000_000

	return (rand.Intn((max - min)) + max)
}
func main() {

	cards := [4]CreditCard{
		{CardNumber: "6037996215467896", ExpireTime: "01/02", Cvv2: 236, BankName: "Melli"},
		{CardNumber: "6004996215467896", ExpireTime: "02/04", Cvv2: 5789, BankName: "Mellat"},
		{CardNumber: "5892996215467896", ExpireTime: "00/07", Cvv2: 187, BankName: "Sapeh"},
		{CardNumber: "5041621546789645", ExpireTime: "04/03", Cvv2: 276, BankName: "Resalt"},
	}

	for _, item := range cards {
		if item.ExpireTime < "01/02" {
			fmt.Println("the card", item.CardNumber, "was expired")
			continue
		}
		remainAmount := getRemainAmountAccount(item.CardNumber, item.ExpireTime)
		fmt.Println("the card", item.CardNumber, "have", remainAmount, "remain amount")
	}
}
