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
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	ctxRoot := context.Background()
	ctx, cancel := context.WithCancel(ctxRoot)
	amqpFactory := amqp.NewAmqpFactory("amqp://admin:admin@rabbitmq:5672", logger)
	go amqpFactory.RunReconnect(ctx, "amqp://admin:admin@rabbitmq:5672")
	pgxCon := postgres.GetConnection(ctx, "postgres://postgre:admin@postgres:5432/servicebalances", logger)

	walletAdapter := postgres.NewWalletAdapter(pgxCon)
	balanceAdapter := postgres.NewBalanceAdapter(pgxCon)

	balanceService := services.NewBalanceService(logger, &walletAdapter, &balanceAdapter)
	amqpFactory.InitRmq()
	walletSender := amqpFactory.NewSender(statics.ExNameBalances, statics.RkGetWalletInfoResponse)
	lockSender := amqpFactory.NewSender(statics.ExNameBalances, statics.RkLockBalanceResponse)
	apiRouter := api.NewApi(logger, &walletSender, &lockSender, &balanceService)
	handler := hanlder.NewHandler(logger, apiRouter)
	listenerEmmitBalance := amqp.NewAmqpListener[balances.EmmitBalanceRequest](ctx, amqpFactory, statics.EmmitBalanceRequestQueue, amqp.GetParserEmmitBalanceRequest(), handler.GetEmmitBalanceHanler())
	listenerGetWalletInfo := amqp.NewAmqpListener[balances.GetWalletInfoRequest](ctx, amqpFactory, statics.GetWalletInfoRequestQueue, amqp.GetParserWalletInfoRequest(), handler.GetWalletInfoHandler())
	listenerLockBalance := amqp.NewAmqpListener[balances.LockBalanceRequest](ctx, amqpFactory, statics.LockBalanceRequestQueue, amqp.GetParserLockBalanceRequest(), handler.GetLockBalanceHandler())
	go listenerEmmitBalance.Run(ctx)
	go listenerGetWalletInfo.Run(ctx)
	go listenerLockBalance.Run(ctx)
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
		default:
			time.Sleep(2 * time.Second)
		}

	}
}
