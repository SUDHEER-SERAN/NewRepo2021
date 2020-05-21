package database

import (
	"context"
	"jobreport/internal/common"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type PostgresDatabase struct {
	conn *pgx.Conn
}

func NewPostgresDatabase(lc fx.Lifecycle, config common.PostgresDatabaseConfig) *PostgresDatabase {
	conn, err := pgx.Connect(context.Background(), config.ConnectionString)

	if err != nil {
		logrus.WithError(err).Fatal("unable to connect to postgres")
	}
	logrus.Info("Connected to Postgres")

	psqldb := &PostgresDatabase{
		conn: conn,
	}

	return psqldb

}

func (p *PostgresDatabase) CreateUserDB() string {
	return "Haisss"
}
