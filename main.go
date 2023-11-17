package main

import (
	"balance-service/api"
	"balance-service/external/balances"
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
	amqpFactory.InitRmq()
	walletSender := amqpFactory.NewSender("e.balances.forward", "r.balances.balance-service.GetWalletInfoResponse.#")
	apiRouter := api.NewApi(logger, &walletSender, &balanceService)
	handler := hanlder.NewHandler(logger, apiRouter)
	listenerEmmitBalance := amqp.NewAmqpListener[balances.EmmitBalanceRequest](ctx, amqpFactory, statics.EmmitBalanceRequestQueue, amqp.GetParserEmmitBalanceRequest(), handler.GetEmmitBalanceHanler())
	listenerGetWalletInfo := amqp.NewAmqpListener[balances.GetWalletInfoRequest](ctx, amqpFactory, statics.GetWalletInfoRequestQueue, amqp.GetParserWalletInfoRequest(), handler.GetWalletInfoHandler())

	go listenerEmmitBalance.Run(ctx)
	go listenerGetWalletInfo.Run(ctx)

	exit := make(chan os.Signal, 1)
	for {
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		select {
		case <-exit:
			{
				cancel()

				pgxCon.Close()
				amqpFactory.ClsConnection()
				break
			}
		}

	}
}
