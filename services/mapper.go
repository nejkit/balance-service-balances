package services

import (
	"balance-service/external/balances"
	"balance-service/sql"
)

func Mapper(walletInfo *sql.WalletModel, balancesInfo []*sql.BalanceModel) balances.WalletInfo {
	if walletInfo == nil {
		return balances.WalletInfo{}
	}
	var balanceProtoModels []*balances.BalanceInfo
	for _, balanceInfo := range balancesInfo {
		balanceProtoModels = append(balanceProtoModels, &balances.BalanceInfo{
			Id:            balanceInfo.Id,
			Currency:      balanceInfo.Currency,
			ActualBalance: balanceInfo.ActualBalance,
			FreezeBalance: balanceInfo.FreezeBalance,
		})
	}

	return balances.WalletInfo{
		Address:      walletInfo.Id,
		Created:      uint64(walletInfo.Created.Unix()),
		BalanceInfos: balanceProtoModels,
	}
}
