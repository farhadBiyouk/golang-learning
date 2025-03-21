package main

import (
	"io"
	"os"
)

func main() {

	defer println("Bye")
	defer println("Good")

	for i := 0; i < 10; i++ {
		defer println(i)
	}

	println("Hello")

	CopyFile("/home/farhad/project/go-learning/8-functions/Defer/destination.txt", "/home/farhad/project/go-learning/8-functions/Defer/source.txt")

}

func CopyFile(destinationName, sourceName string) (written int64, err error) {
	source, err := os.Open(sourceName)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return
	}

	defer destination.Close()

	return io.Copy(destination, source)
}
