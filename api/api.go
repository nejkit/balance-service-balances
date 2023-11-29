package api

import (
	"balance-service/abstractions"
	"balance-service/external/balances"
	"context"

	logger "github.com/sirupsen/logrus"
)

type BalanceApi struct {
	balanceservice   abstractions.BalanceService
	walletInfoSender abstractions.AmqpSender
	lockInfoSender   abstractions.AmqpSender
}

func NewApi(walletInfoSender abstractions.AmqpSender, lockInfoSender abstractions.AmqpSender, bs abstractions.BalanceService) BalanceApi {
	return BalanceApi{walletInfoSender: walletInfoSender, balanceservice: bs, lockInfoSender: lockInfoSender}
}

func (r *BalanceApi) EmmitBalanceApi(ctx context.Context, request *balances.EmmitBalanceRequest) {
	go r.balanceservice.EmmitBalance(ctx, request)
}

func (r *BalanceApi) GetWalletInfoApi(ctx context.Context, request *balances.GetWalletInfoRequest) {
	response := r.balanceservice.GetInfoAboutBalance(ctx, request)
	logger.Infoln("Response body: ", response.String())
	go r.walletInfoSender.SendMessage(ctx, response)
}

func (r *BalanceApi) LockBalanceApi(ctx context.Context, request *balances.LockBalanceRequest) {
	response := r.balanceservice.LockBalance(ctx, request)
	logger.Infoln("Response body: ", response.String())
	go r.lockInfoSender.SendMessage(ctx, response)
}

func (r *BalanceApi) TransferApi(ctx context.Context, request *balances.CreateTransferRequest) {
	go r.balanceservice.ProcessTransfer(ctx, request)
}
