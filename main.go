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
	msgsEmmit := rmq.InitListener("q.balances.request.EmmitBalanceRequest")
	msgsGet := rmq.InitListener("q.balances.request.GetWalletInfoRequest")
	forever := make(chan (bool))
	go func() {
		for msg := range msgsEmmit {
			var request proto.EmmitBalanceRequest
			err := protoTuls.Unmarshal(msg.Body, &request)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Received Request: Id: %s, Address: %s, Amount: %d, Currency: %s\n",
				request.GetId(), request.GetAddress(), request.GetAmount(), request.GetCurrency())
			err = db.EmmitBalance(&request)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
	}()
	go func() {
		for msg := range msgsGet {
			var request proto.GetWalletInfoRequest
			err := protoTuls.Unmarshal(msg.Body, &request)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Received request for GetWalletInfo. Body{ Id: %s, Address: %s}", request.GetId(), request.GetAddress())
			fmt.Println()
			data := db.GetInfoAboutBalance(request.Address)
			response := proto.GetWalletInfoResponse{
				Id:         request.Id,
				WalletInfo: data,
			}
			rmq.SendResponseGetWalletInfo(&response)
			fmt.Println("Response was sended")
		}
	}()

	<-forever

}
