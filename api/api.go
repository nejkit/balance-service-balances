package api

import (
	"balance-service/abstractions"
	"balance-service/external/balances"
	"context"

	"github.com/sirupsen/logrus"
)

type BalanceApi struct {
	logger           *logrus.Logger
	balanceservice   abstractions.BalanceService
	walletInfoSender abstractions.AmqpSender
	lockInfoSender   abstractions.AmqpSender
}

func NewApi(logger *logrus.Logger, walletInfoSender abstractions.AmqpSender, lockInfoSender abstractions.AmqpSender, bs abstractions.BalanceService) BalanceApi {
	return BalanceApi{logger: logger, walletInfoSender: walletInfoSender, balanceservice: bs, lockInfoSender: lockInfoSender}
}

func (r *BalanceApi) EmmitBalanceApi(ctx context.Context, request *balances.EmmitBalanceRequest) {
	r.balanceservice.EmmitBalance(ctx, request)
}

func (r *BalanceApi) GetWalletInfoApi(ctx context.Context, request *balances.GetWalletInfoRequest) {
	response := r.balanceservice.GetInfoAboutBalance(ctx, request)
	r.logger.Infoln("Response body: ", response.String())
	go r.walletInfoSender.SendMessage(ctx, response)
}

func (r *BalanceApi) LockBalanceApi(ctx context.Context, request *balances.LockBalanceRequest) {
	response := r.balanceservice.LockBalance(ctx, request)
	r.logger.Infoln("Response body: ", response.String())
	go r.lockInfoSender.SendMessage(ctx, response)
}
