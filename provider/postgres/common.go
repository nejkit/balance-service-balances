package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func GetConnection(ctx context.Context, connectionString string, logger *logrus.Logger) *pgx.Conn {
	var con *pgx.Conn
	var err error
	for {
		con, err = pgx.Connect(ctx, connectionString)
		if err != nil {
			logger.Errorln("Connection failed! ", err.Error())
			continue
		}
		break
	}
	return con
}
