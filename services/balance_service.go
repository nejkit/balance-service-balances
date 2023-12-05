package services

import (
	"balance-service/abstractions"
	"balance-service/external/balances"
	"context"

	logger "github.com/sirupsen/logrus"
)

type BalanceService struct {
	walletAdapter  abstractions.WalletAdapter
	balanceAdapter abstractions.BalanceAdapter
	transferSender abstractions.AmqpSender
}

func NewBalanceService(walletAdapter abstractions.WalletAdapter, balanceAdapter abstractions.BalanceAdapter, ts abstractions.AmqpSender) BalanceService {
	return BalanceService{walletAdapter: walletAdapter, balanceAdapter: balanceAdapter, transferSender: ts}
}

func (s *BalanceService) EmmitBalance(ctx context.Context, request *balances.EmmitBalanceRequest) {
	if request.GetCurrency() == "" {
		logger.Errorln("Wrong currency, skip event")
		return
	}

	if request.GetAmount() <= 0 {
		logger.Errorln("Wrong amount, skip event")
		return
	}

	if request.GetAddress() == "" {
		logger.Errorln("Wrong address, skip event")
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
	logger.Infoln("Id: ", balanceInfo.Id, "Amount: ", balanceInfo.ActualBalance)
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

func (s *BalanceService) UnLockBalance(ctx context.Context, request *balances.UnLockBalanceRequest) {
	go s.balanceAdapter.UnlockTransferBalance(ctx, request.GetAddress(), float64(request.GetAmount()), request.GetCurrency())
}

func (s *BalanceService) ProcessTransfer(ctx context.Context, request *balances.CreateTransferRequest) {
	response := &balances.Transfer{
		Id:            request.Id,
		SenderData:    request.SenderData,
		RecepientData: request.RecepientData,
		State:         balances.TransferState_TRANSFER_STATE_IN_PROGRESS,
	}
	go s.transferSender.SendMessage(ctx, response)
	balanceSenderCur1 := s.balanceAdapter.GetBalanceInfo(ctx, request.GetSenderData().GetAddress(), request.GetSenderData().GetCurrency())
	balanceSenderCur2 := s.balanceAdapter.GetBalanceInfo(ctx, request.GetRecepientData().GetAddress(), request.GetRecepientData().GetCurrency())
	balanceRecepientCur1 := s.balanceAdapter.GetBalanceInfo(ctx, request.GetRecepientData().GetAddress(), request.GetSenderData().GetCurrency())
	balanceRecepientCur2 := s.balanceAdapter.GetBalanceInfo(ctx, request.GetSenderData().GetAddress(), request.GetRecepientData().GetCurrency())
	if balanceSenderCur1.FreezeBalance < float64(request.SenderData.GetAmount()) {
		s.rejectTransfer(ctx, response)
		return
	}
	if balanceSenderCur2.FreezeBalance < float64(request.RecepientData.GetAmount()) {
		s.rejectTransfer(ctx, response)
		return
	}

	if err := s.balanceAdapter.TransferMoney(
		ctx,
		balanceSenderCur1.Id,
		balanceSenderCur2.Id,
		balanceRecepientCur1.Id,
		balanceRecepientCur2.Id,
		request.SenderData.Amount,
		request.RecepientData.Amount,
	); err != nil {
		response.Error = &balances.BalanceErrorMessage{
			ErrorCode: balances.BalancesErrorCodes_BALANCE_ERROR_CODE_INTERNAL,
			Message:   "InternalError",
		}
		response.State = balances.TransferState_TRANSFER_STATE_REJECT
		s.transferSender.SendMessage(ctx, response)
		return
	}

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
