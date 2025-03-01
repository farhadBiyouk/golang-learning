package main

import (
	"LogToFile/core"

	"github.com/pkg/errors"
)

type Person struct {
	Name string
	Age  int
}

var logger = core.NewFileLogger()

func main() {

	p1 := &Person{Name: "ali", Age: 23}
	logger.Info().Interface("Body", p1).Str("Category", "Search").
		Int("DurationTime", 80).
		Msg("Searching for a thing")

	err := func3()
	if err != nil {
		logger.Error().Stack().Err(err).Msg("this is a stack error")
	}
}

func func1() error {
	return errors.New("this is an error (func1)")
}

func func2() error {
	err := func1()
	if err != nil {
		return err
	}
	return nil
}

func func3() error {
	err := func2()
	if err != nil {
		return err
	}
	return nil
}
