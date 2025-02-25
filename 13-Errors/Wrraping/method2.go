package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type IoError struct {
	FileName string
	Message  string
	Err      error
}

func (ioe *IoError) Error() string {
	return fmt.Sprintf("IO error occurred: FileName: %s Message: %s Detail: %s", ioe.FileName, ioe.Message, ioe.Err.Error())
}

func (ior *IoError) Unwrap() error {
	return ior.Err
}

func main() {
	err := CopyFile1("/home/farhad/project/go-learning/13-Errors/Wrraping/sourcce.txt", "dest.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Error: %s\n", errors.Unwrap(err))
	}
}

func CopyFile1(sourceName, destinationName string) error {

	source, err := os.Open(sourceName)

	if err != nil {
		return &IoError{FileName: sourceName, Message: "during copy file could not open source file", Err: err}
	}

	defer source.Close()

	destination, err := os.Create(destinationName)

	if err != nil {
		return &IoError{FileName: destinationName, Message: "during copy file could not create destination file", Err: err}
	}

	defer destination.Close()

	_, err = io.Copy(destination, source)

	if err != nil {
		return &IoError{FileName: sourceName, Message: "during copy file could not copy file", Err: err}
	}

	return nil

}
