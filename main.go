package main

import (
	"balance-service/api"
	rmq "balance-service/rmq"
	"balance-service/services"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	chanRabbit := rmq.InitRmq("amqp://admin:admin@rabbitmq:5672")
	pgxConnection := services.InitConnection("postgres://postgre:admin@postgres:5432")
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	var vg sync.WaitGroup
	vg.Add(3)
	go api.InnitEmmitBalanceApi(chanRabbit, pgxConnection, logger)
	go api.InnitGetWalletInfoApi(chanRabbit, pgxConnection, logger)
	vg.Wait()
}
