package services

import (
	"balance-service/abstractions"
	"balance-service/external/balances"
	"context"

	"github.com/sirupsen/logrus"
)

type BalanceService struct {
	logger         *logrus.Logger
	walletAdapter  abstractions.WalletAdapter
	balanceAdapter abstractions.BalanceAdapter
}

func NewBalanceService(logger *logrus.Logger, walletAdapter abstractions.WalletAdapter, balanceAdapter abstractions.BalanceAdapter) BalanceService {
	return BalanceService{walletAdapter: walletAdapter, balanceAdapter: balanceAdapter, logger: logger}
}

func (s *BalanceService) EmmitBalance(ctx context.Context, request *balances.EmmitBalanceRequest) {
	if request.GetCurrency() == "" {
		s.logger.Errorln("Wrong currency, skip event")
		return
	}

	if request.GetAmount() <= 0 {
		s.logger.Errorln("Wrong amount, skip event")
		return
	}

	if request.GetAddress() == "" {
		s.logger.Errorln("Wrong address, skip event")
		return
	}

	walletInfo := s.walletAdapter.GetWalletInfo(ctx, request.GetAddress())
	if walletInfo == nil {
		s.walletAdapter.InsertWalletInfo(ctx, request)
	}

	balanceInfo := s.balanceAdapter.GetBalanceInfo(ctx, request.GetAddress(), request.GetCurrency())
	if balanceInfo == nil {
		go s.balanceAdapter.InsertBalanceInfo(ctx, request)
	} else {
		go s.balanceAdapter.EmmitBalanceInfo(ctx, request, balanceInfo.Id)
	}
}

func (s *BalanceService) GetInfoAboutBalance(ctx context.Context, request *balances.GetWalletInfoRequest) *balances.GetWalletInfoResponse {
	walletInfo := s.walletAdapter.GetWalletInfo(ctx, request.GetAddress())
	balancesInfo := s.balanceAdapter.GetBalancesInfo(ctx, request.GetAddress())
	responseBody := Mapper(walletInfo, balancesInfo)
	return &balances.GetWalletInfoResponse{Id: request.GetId(), WalletInfo: &responseBody}
}

func (s *BalanceService) LockBalance(ctx context.Context, request *balances.LockBalanceRequest) *balances.LockBalanceResponse {
	balanceInfo := s.balanceAdapter.GetBalanceInfo(ctx, request.GetAddress(), request.GetCurrency())
	if balanceInfo == nil {
		return &balances.LockBalanceResponse{
			Id:    request.GetId(),
			State: balances.LockBalanceStatus_REJECTED,
			ErrorMessage: &balances.ErrorMessage{
				ErrorCode: balances.ErrorCodes_ERROR_CODE_NOT_EXISTS_BALANCE,
				Message:   "Balance not exists",
			},
		}
	}
	s.logger.Infoln("Id: ", balanceInfo.Id, "Amount: ", balanceInfo.ActualBalance)
	if balanceInfo.ActualBalance < float64(request.GetAmount()) {
		return &balances.LockBalanceResponse{
			Id:    request.GetId(),
			State: balances.LockBalanceStatus_REJECTED,
			ErrorMessage: &balances.ErrorMessage{
				ErrorCode: balances.ErrorCodes_ERROR_CODE_NOT_ENOUGH_BALANCE,
				Message:   "Not enough balance",
			},
		}
	}

	s.balanceAdapter.LockTransferBalance(ctx, balanceInfo.Id, float64(request.GetAmount()))
	return &balances.LockBalanceResponse{
		Id:    request.GetId(),
		State: balances.LockBalanceStatus_DONE,
	}
}
