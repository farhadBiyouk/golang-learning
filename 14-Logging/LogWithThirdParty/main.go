package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// 	log.Print("test")

	// 	log.Info().Str("Category", "Search").Int("DurationTime", 80).Msg("Searching for any thing")

	// 	err := errors.New("Error occurred")
	// 	log.Error().Err(err).Msg("This is a error ")
	err := func3()
	if err != nil {
		log.Error().Stack().Err(err).Msg("this is a stack error")
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
