package user

import (
	"context"
	"errors"
	"jobreport/internal/database"

	"github.com/sirupsen/logrus"
)

type Database struct {
	database *database.PostgresDatabase
}

func NewUserDatabase(database *database.PostgresDatabase) *Database {
	return &Database{
		database: database,
	}
}
func (d *Database) CreateUserRepo(ctx context.Context, u User) error {
	if _, err := d.database.Conn.
		Exec(ctx, "insert into users(username,password,address,role) values($1,$2,$3,$4)", u.Username, u.Password, u.Address, u.Role); err != nil {
		logrus.WithError(err).Warn("unable to insert doc")
		return errors.New("unable to insert doc metadata")
	}
	return nil
}

func (d *Database) FindUser(user User) error {
	return nil
}
