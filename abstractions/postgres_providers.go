package abstractions

import (
	"balance-service/sql"
	"context"

	"github.com/nejkit/processing-proto/balances"
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
}
