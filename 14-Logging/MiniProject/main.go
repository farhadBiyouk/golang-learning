package main

import (
	"notification/core"
	"notification/entities"
	"notification/services"
)

var logger = core.NewFileLogger()

func main() {

	order1 := entities.Order{
		ID:               1,
		UserFullName:     "John Doe",
		UserId:           "09115252524",
		Price:            float64(100),
		Status:           true,
		NotificationType: entities.Email,
	}
	order2 := entities.Order{
		ID:               2,
		UserFullName:     "John Mac",
		UserId:           "09123332221",
		Price:            float64(220),
		Status:           true,
		NotificationType: entities.Sms,
	}

	order3 := entities.Order{
		ID:               3,
		UserFullName:     "Reza Mac",
		UserId:           "09991252525",
		Price:            float64(210),
		Status:           true,
		NotificationType: "",
	}

	orderService := services.NewOrderService()
	_, err := orderService.CreateOrder(&order1)
	if err != nil {
		logger.Error().Interface("entity info", order1).Err(err).Msg("order 1 is not valid`")
	}

	_, err = orderService.CreateOrder(&order2)
	if err != nil {
		logger.Error().Interface("entity info", order2).Err(err).Msg("order 1 is not valid`")
	}

	_, err = orderService.CreateOrder(&order3)
	if err != nil {
		logger.Error().Interface("entity info", order3).Err(err).Msg("order 1 is not valid`")
	}
}
