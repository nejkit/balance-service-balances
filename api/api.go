package api

import (
	"balance-service/abstractions"
	"balance-service/hanlder"
	"context"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type RouterApi struct {
	logger            *logrus.Logger
	handler           hanlder.Handler
	emmitBalanceRoute <-chan amqp091.Delivery
	walletInfoRoute   <-chan amqp091.Delivery
	walletInfoSender  abstractions.AmqpSender
}

func NewRouter(logger *logrus.Logger, handler hanlder.Handler, emBalRoute <-chan amqp091.Delivery, walInfRoute <-chan amqp091.Delivery, walletInfoSender abstractions.AmqpSender) RouterApi {
	return RouterApi{logger: logger, handler: handler, emmitBalanceRoute: emBalRoute, walletInfoRoute: walInfRoute, walletInfoSender: walletInfoSender}
}

func (r *RouterApi) StartApi(ctx context.Context) {
	for {
		select {
		case <-r.emmitBalanceRoute:
			{
				r.handler.EmmitBalance(ctx, <-r.emmitBalanceRoute)
			}
		case <-r.walletInfoRoute:
			{
				response := r.handler.GetWalletInfo(ctx, <-r.walletInfoRoute)
				r.walletInfoSender.SendMessage(ctx, response)
			}
		}
	}
}
