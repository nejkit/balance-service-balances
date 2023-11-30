package abstractions

import (
	"balance-service/external/balances"
	"context"
)

type BalanceService interface {
	EmmitBalance(ctx context.Context, request *balances.EmmitBalanceRequest)
	GetInfoAboutBalance(ctx context.Context, request *balances.GetWalletInfoRequest) *balances.GetWalletInfoResponse
	LockBalance(ctx context.Context, request *balances.LockBalanceRequest) *balances.LockBalanceResponse
	UnLockBalance(ctx context.Context, request *balances.UnLockBalanceRequest)
	ProcessTransfer(ctx context.Context, request *balances.CreateTransferRequest)
}
