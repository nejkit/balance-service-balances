package main

import (
	rmq "balance-service/rmq"
	"balance-service/services"
)

func main() {
	chanRabbit := rmq.InitRmq("")
	msgsEmmit := rmq.InitListener("q.balances.request.EmmitBalanceRequest", chanRabbit)
	msgsGet := rmq.InitListener("q.balances.request.GetWalletInfoRequest", chanRabbit)

	postgreConnection := services.InitConnection("")

	go services.EmmitBalance()
	go services.GetInfoAboutBalance()

}
