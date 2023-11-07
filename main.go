package main

import (
	rmq "balance-service/rmq"
	db "balance-service/sql"
	"fmt"

	proto "github.com/nejkit/processing-proto/balances"
	protoTuls "google.golang.org/protobuf/proto"
)

func main() {
	rmq.InitRmq()
	msgs := rmq.InitListener()
	forever := make(chan (bool))
	conDb := db.InitDb()
	go func() {
		for msg := range msgs {
			var request proto.EmmitBalanceRequest
			err := protoTuls.Unmarshal(msg.Body, &request)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Received Request: Id: %s, Address: %s, Amount: %d, Currency: %s\n",
				request.GetId(), request.GetAddress(), request.GetAmount(), request.GetCurrency())
			err = db.EmmitBalance(&request, conDb)
			fmt.Println()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	<-forever

}
