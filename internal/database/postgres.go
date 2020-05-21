package database

import (
	"context"
	"jobreport/internal/common"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type PostgresDatabase struct {
	Conn *pgx.Conn
}

func NewPostgresDatabase(lc fx.Lifecycle, config common.PostgresDatabaseConfig) *PostgresDatabase {
	conn, err := pgx.Connect(context.Background(), config.ConnectionString)

	if err != nil {
		logrus.WithError(err).Fatal("unable to connect to postgres")
	}
	logrus.Info("Connected to Postgres")

	psqldb := &PostgresDatabase{
		Conn: conn,
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logrus.Info("CLosing the DB Connection")
			return conn.Close(context.Background())
		},
	})
	return psqldb

}
