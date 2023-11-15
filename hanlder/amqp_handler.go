package hanlder

import (
	"balance-service/abstractions"
	"context"

	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	logger         *logrus.Logger
	balanceService abstractions.BalanceService
}

func NewHandler(logger *logrus.Logger, balanceService abstractions.BalanceService) Handler {
	return Handler{logger: logger, balanceService: balanceService}
}

func (h *Handler) EmmitBalance(ctx context.Context, msg amqp091.Delivery) {
	var request balances.EmmitBalanceRequest
	err := proto.Unmarshal(msg.Body, &request)
	if err != nil {
		h.logger.Errorln("Error parse message. Skipping...")
		msg.Nack(false, false)
	}
	msg.Ack(false)
	h.balanceService.EmmitBalance(ctx, &request)
}

func (h *Handler) GetWalletInfo(ctx context.Context, msg amqp091.Delivery) *balances.GetWalletInfoResponse {
	var request balances.GetWalletInfoRequest
	err := proto.Unmarshal(msg.Body, &request)
	if err != nil {
		h.logger.Errorln("Error parse message. Skipping...")
		msg.Nack(false, false)
	}
	msg.Ack(false)
	return h.balanceService.GetInfoAboutBalance(ctx, &request)

}
