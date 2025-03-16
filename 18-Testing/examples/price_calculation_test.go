package examples

import "testing"

func TestPriceCalCulator(t *testing.T) {
	//Arrange
	expectPrice := 1526000
	nights := 2
	personCont := 2
	//Act
	actual := CalculateRoomPrice("Single", nights, personCont)

	//Assert
	if actual != expectPrice {
		t.Errorf("ex %v - ac %v ", expectPrice, actual)
	}

}
