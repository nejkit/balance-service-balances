package hanlder

import (
	"balance-service/api"
	"balance-service/external/balances"
	"context"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Logger
	api    api.BalanceApi
}

func NewHandler(logger *logrus.Logger, balanceapi api.BalanceApi) Handler {
	return Handler{logger: logger, api: balanceapi}
}

func (h *Handler) GetEmmitBalanceHanler() func(context.Context, *balances.EmmitBalanceRequest) {
	return func(ctx context.Context, ebr *balances.EmmitBalanceRequest) {
		h.api.EmmitBalanceApi(ctx, ebr)
	}
}

func (h *Handler) GetWalletInfoHandler() func(context.Context, *balances.GetWalletInfoRequest) {
	return func(ctx context.Context, gwir *balances.GetWalletInfoRequest) {
		h.logger.Infoln("Event body: ", gwir.String())
		h.api.GetWalletInfoApi(ctx, gwir)
	}
}
