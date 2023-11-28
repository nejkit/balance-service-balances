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
	transferSender abstractions.AmqpSender
}

func NewBalanceService(logger *logrus.Logger, walletAdapter abstractions.WalletAdapter, balanceAdapter abstractions.BalanceAdapter, ts abstractions.AmqpSender) BalanceService {
	return BalanceService{walletAdapter: walletAdapter, balanceAdapter: balanceAdapter, logger: logger, transferSender: ts}
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
			ErrorMessage: &balances.BalanceErrorMessage{
				ErrorCode: balances.BalancesErrorCodes_BALANCE_ERROR_CODE_NOT_EXISTS_BALANCE,
				Message:   "Balance not exists",
			},
		}
	}
	s.logger.Infoln("Id: ", balanceInfo.Id, "Amount: ", balanceInfo.ActualBalance)
	if balanceInfo.ActualBalance < float64(request.GetAmount()) {
		return &balances.LockBalanceResponse{
			Id:    request.GetId(),
			State: balances.LockBalanceStatus_REJECTED,
			ErrorMessage: &balances.BalanceErrorMessage{
				ErrorCode: balances.BalancesErrorCodes_BALANCE_ERROR_CODE_NOT_ENOUGH_BALANCE,
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

func (s *BalanceService) ProcessTransfer(ctx context.Context, request *balances.CreateTransferRequest) {
	response := &balances.Transfer{
		Id:            request.Id,
		SenderData:    request.SenderData,
		RecepientData: request.RecepientData,
		State:         balances.TransferState_TRANSFER_STATE_IN_PROGRESS,
	}
	go s.transferSender.SendMessage(ctx, response)
	balanceSender := s.balanceAdapter.GetBalanceInfo(ctx, request.GetSenderData().GetAddress(), request.GetSenderData().GetCurrency())
	balanceRecepient := s.balanceAdapter.GetBalanceInfo(ctx, request.GetRecepientData().GetAddress(), request.GetRecepientData().GetCurrency())
	if balanceSender.FreezeBalance < float64(request.SenderData.GetAmount()) {
		s.rejectTransfer(ctx, response)
		return
	}
	if balanceRecepient.FreezeBalance < float64(request.RecepientData.GetAmount()) {
		s.rejectTransfer(ctx, response)
		return
	}
	s.balanceAdapter.TransferMoney(ctx, *balanceSender, *balanceRecepient, float64(request.SenderData.Amount))
	s.balanceAdapter.TransferMoney(ctx, *balanceRecepient, *balanceSender, float64(request.RecepientData.Amount))

	response.State = balances.TransferState_TRANSFER_STATE_DONE
	s.transferSender.SendMessage(ctx, response)
}

func (s *BalanceService) rejectTransfer(ctx context.Context, response *balances.Transfer) {

	response.State = balances.TransferState_TRANSFER_STATE_REJECT
	response.Error = &balances.BalanceErrorMessage{
		ErrorCode: balances.BalancesErrorCodes_BALANCE_ERROR_CODE_NOT_ENOUGH_BALANCE,
		Message:   "Not enough balance",
	}
	go s.transferSender.SendMessage(ctx, response)
	return
}
