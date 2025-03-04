package examples

import "fmt"

func Intro() {
	numChannel := make(chan int, 0)

	go func() {
		numChannel <- 5
	}()
	receivedNum := <-numChannel

	fmt.Println(receivedNum)
}
