package abstractions

import (
	"context"

	"github.com/nejkit/processing-proto/balances"
)

type BalanceService interface {
	EmmitBalance(ctx context.Context, request *balances.EmmitBalanceRequest)
	GetInfoAboutBalance(ctx context.Context, request *balances.GetWalletInfoRequest) *balances.GetWalletInfoResponse
}
