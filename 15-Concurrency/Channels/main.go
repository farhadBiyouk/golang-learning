package main

import (
	"fmt"
	"time"
)

func main() {
	//define channel

	Ch1 := make(chan int) // this is an unbuffered channel, so needed receiver
	fmt.Println("sent data to channel")
	go SendDataToChannel(10, Ch1)

	fmt.Println("received data from channel")
	go ReceiveDataFromChannel(Ch1)

	time.Sleep(time.Second * 2 )

}

func SendDataToChannel(number int, ch chan int) {
	ch <- number
}

func ReceiveDataFromChannel(ch chan int) {
	value := <-ch
	fmt.Println(value)
}
