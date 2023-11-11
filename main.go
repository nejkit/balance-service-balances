package main

import (
	"balance-service/api"
	rmq "balance-service/rmq"
	"balance-service/services"

	"github.com/sirupsen/logrus"
)

func main() {
	chanRabbit := rmq.InitRmq("amqp://admin:admin@rabbitmq:5672")
	pgxConnection := services.InitConnection("postgres://postgre:admin@postgres:5432")
	var logger logrus.Logger
	logger.SetLevel(logrus.InfoLevel)
	api.InnitEmmitBalanceApi(chanRabbit, pgxConnection, &logger)
	api.InnitGetWalletInfoApi(chanRabbit, pgxConnection, &logger)
}
