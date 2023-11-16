package main

import (
	"balance-service/api"
	"balance-service/hanlder"
	"balance-service/provider/amqp"
	"balance-service/provider/postgres"
	"balance-service/services"
	"balance-service/statics"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	ctxRoot := context.Background()
	ctx, cancel := context.WithCancel(ctxRoot)
	amqpFactory := amqp.NewAmqpFactory("amqp://admin:admin@rabbitmq:5672", logger)
	pgxCon := postgres.GetConnection(ctx, "postgres://postgre:admin@postgres:5432/servicebalances", logger)
	walletAdapter := postgres.NewWalletAdapter(pgxCon)
	balanceAdapter := postgres.NewBalanceAdapter(pgxCon)
	balanceService := services.NewBalanceService(logger, &walletAdapter, &balanceAdapter)
	handler := hanlder.NewHandler(logger, &balanceService)

	emmitBalanceChannel := amqpFactory.NewRmqChan()
	walletInfoListenerChannel := amqpFactory.NewRmqChan()

	emmitBalanceRoute := amqpFactory.NewConsumeChan(statics.EmmitBalanceRequestQueue, emmitBalanceChannel)
	walletInfoRoute := amqpFactory.NewConsumeChan(statics.GetWalletInfoRequestQueue, walletInfoListenerChannel)
	walletInfoSender := amqpFactory.NewSender("e.balances.forward", "r.balances.#.GetWalletInfoResponse.#")

	apiRouter := api.NewRouter(logger, handler, emmitBalanceRoute, walletInfoRoute, &walletInfoSender)
	go apiRouter.StartApi(ctx)

	exit := make(chan os.Signal, 1)
	for {
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		select {
		case <-exit:
			{
				cancel()

				emmitBalanceChannel.Close()
				walletInfoListenerChannel.Close()
				pgxCon.Close()
				amqpFactory.ClsConnection()
				break
			}
		}
	}
}
