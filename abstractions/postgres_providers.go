package abstractions

import (
	"balance-service/external/balances"
	"balance-service/sql"
	"context"
)

type WalletAdapter interface {
	InsertWalletInfo(ctx context.Context, info *balances.EmmitBalanceRequest)
	GetWalletInfo(ctx context.Context, address string) *sql.WalletModel
}

type BalanceAdapter interface {
	EmmitBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest, id string)
	InsertBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest)
	GetBalancesInfo(ctx context.Context, address string) []*sql.BalanceModel
	GetBalanceInfo(ctx context.Context, address string, currency string) *sql.BalanceModel
	LockTransferBalance(ctx context.Context, id string, amount float64)
	UnlockTransferBalance(ctx context.Context, address string, amount float64, cur string)
	TransferMoney(ctx context.Context,
		sender1cur, sender2cur, recepient1cur, recepient2cur string, senderAmount, recepientAmount float64) error
}
