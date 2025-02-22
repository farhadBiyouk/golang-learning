package main

import "fmt"

type Test struct {
	Id   int
	name string
}

func main() {
	test := Test{Id: 1, name: "test"}
	test.Print()
	test.print2()
}

func (test Test) Print() { // public type
	fmt.Println(test.Id, test.name)
}
func (test Test) print2() { // private type
	fmt.Println(test.Id, test.name)
}

// this is  Facade pattern

func AddToShoppingCart(productId int, userId int, quantity int) { // public type
	checkUserStatus(userId)
	checkProductStatus(productId)
	checkQuantity(productId, quantity)
	checkPriceChanges(productId)
	addToCart(productId, quantity)
	updateShoppingCartTotalPrice()
}

// all function private type

func checkUserStatus(userId int) {
	// check user status
}

func checkProductStatus(productId int) {
	// check product status
}

func checkQuantity(productId int, quantity int) {
	// check product quantity
}

func checkPriceChanges(productId int) {
	// check price changes
}

func addToCart(productId int, quantity int) {
	// add to cart
}

func updateShoppingCartTotalPrice() {
	// update shopping cart total price
}
