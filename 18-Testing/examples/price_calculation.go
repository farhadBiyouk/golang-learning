package examples

func CalculateRoomPrice(room string, nights int, personCount int) (finalPrice int) {
	price := 0
	switch room {
	case "Single":
		price = nights * personCount * 350000
	case "Double":
		price = nights * personCount * 220000
	case "Standard":
		price = nights * personCount * 450000
	case "Suite":
		price = nights * personCount * 1100000
	}

	tax := float64(price) * 0.09
	finalPrice = price + int(tax)

	return
}
