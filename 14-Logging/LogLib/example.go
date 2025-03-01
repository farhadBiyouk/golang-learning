package main

import (
	"fmt"
	"log"
	"os"
)

var (
	InfoLogging  *log.Logger
	WarnLogging  *log.Logger
	ErrorLogging *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	flags := log.Ldate | log.Ltime | log.Lshortfile
	log.SetFlags(flags)
	log.SetOutput(file)

	InfoLogging = log.New(file, "Info: ", flags)
	WarnLogging = log.New(file, "Warning: ", flags)
	ErrorLogging = log.New(file, "Error: ", flags)
}
func main() {
	InfoLogging.Println("start sum function")
	sum(10, 6)
	InfoLogging.Println("end sum function")

}

func sum(a, b int) {
	fmt.Println(a + b)
}
