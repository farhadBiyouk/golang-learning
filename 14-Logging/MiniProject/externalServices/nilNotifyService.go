package externalServices

import (
	"errors"
	"fmt"
	"notification/core"

	"github.com/rs/zerolog"
)

var logger = core.NewFileLogger()

type NilNotifyService struct {
}

func (e *NilNotifyService) SendNotify(receiver string, message string) error {

	if receiver == "" {
		return errors.New("receiver must be not empty value")
	}
	fmt.Printf("nilNotifyService: receiver: %s, message: %s\n", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("Message sent")
	return nil
}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
