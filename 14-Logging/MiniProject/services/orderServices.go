package services

import (
	"errors"
	"fmt"
	"notification/core"
	"notification/entities"
	"notification/externalServices"
)

var logger = core.NewFileLogger()

type OrderService struct {
	externalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) (*entities.Order, error) {
	if !order.Status {
		return order, errors.New("order is not valid")
	}

	if order.Price < 150 {
		return order, errors.New("order price is invalid")
	}

	fmt.Printf("Order created: %v\n", order)
	logger.Info().Interface("order", order).Msg("created order")

	orderService.Notifier = externalServices.NewNotifier(order.NotificationType)

	logger.Info().Msgf("Notifier founded %T", orderService.Notifier)

	err := orderService.SendNotify(order.UserId, "Order created")

	return order, err
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
