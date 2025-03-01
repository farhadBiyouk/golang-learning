package core

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewFileLogger() *zerolog.Logger {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	file, err := os.OpenFile("loging.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("have error in operation")
	}
	logger := zerolog.
		New(file).
		With().
		Timestamp().
		Str("Core app", "in searching").
		Logger()

	return &logger
}
