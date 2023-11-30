package hanlder

import (
	"balance-service/api"
	"balance-service/external/balances"
	"context"

	logger "github.com/sirupsen/logrus"
)

type Handler struct {
	api api.BalanceApi
}

func NewHandler(balanceapi api.BalanceApi) Handler {
	return Handler{api: balanceapi}
}

func (h *Handler) GetEmmitBalanceHanler() func(context.Context, *balances.EmmitBalanceRequest) {
	return func(ctx context.Context, ebr *balances.EmmitBalanceRequest) {
		h.api.EmmitBalanceApi(ctx, ebr)
	}
}

func (h *Handler) GetWalletInfoHandler() func(context.Context, *balances.GetWalletInfoRequest) {
	return func(ctx context.Context, gwir *balances.GetWalletInfoRequest) {
		logger.Infoln("Event body: ", gwir.String())
		h.api.GetWalletInfoApi(ctx, gwir)
	}
}

func (h *Handler) GetLockBalanceHandler() func(context.Context, *balances.LockBalanceRequest) {
	return func(ctx context.Context, lbr *balances.LockBalanceRequest) {
		logger.Infoln("Event body: ", lbr.String())
		h.api.LockBalanceApi(ctx, lbr)
	}
}

func (h *Handler) GetUnLockBalanceHandler() func(context.Context, *balances.UnLockBalanceRequest) {
	return func(ctx context.Context, lbr *balances.UnLockBalanceRequest) {
		logger.Infoln("Event body: ", lbr.String())
		h.api.UnLockBalanceApi(ctx, lbr)
	}
}

func (h *Handler) GetTransferHandler() func(context.Context, *balances.CreateTransferRequest) {
	return func(ctx context.Context, ctr *balances.CreateTransferRequest) {
		logger.Infoln("Event body: ", ctr.String())
		h.api.TransferApi(ctx, ctr)
	}
}
