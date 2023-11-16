package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetConnection(ctx context.Context, connectionString string, logger *logrus.Logger) *pgxpool.Pool {
	var con *pgxpool.Pool
	var err error
	for {
		con, err = pgxpool.New(ctx, connectionString)
		if err != nil {
			logger.Errorln("Connection failed! ", err.Error())
			continue
		}
		break
	}
	return con
}
