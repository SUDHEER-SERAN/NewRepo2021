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
func (d *Database) CreateUserRepo(ctx context.Context, u User, e EmployeeDetails) error {

	var empId int
	if err := d.database.Conn.
		QueryRow(ctx, "insert into temployeedetails(firstname,lastname,mobileno,address,role) values($1,$2,$3,$4,$5) returning empid", e.FirstName, e.LastName, e.MobileNo, e.Address, e.Role).Scan(&empId); err != nil {
		logrus.WithError(err).Warn("unable to insert doc in temployeedetails")
		return errors.New("unable to insert doc metadata in temployeedetails")
	}

	if _, err := d.database.Conn.
		Exec(ctx, "insert into tusers(empid,password,username) values($1,$2,$3)", empId, u.Password, u.Username); err != nil {
		logrus.WithError(err).Warn("unable to insert doc in tusers")
		return errors.New("unable to insert doc metadata in tusers")
	}

	return nil
}

func (d *Database) FindUser(tx context.Context, userDetails User) (error, *User) {
	var user = &User{}
	err := d.database.Conn.
		QueryRow(context.Background(), "select password,empid from tusers where username=$1", userDetails.Username).Scan(&user.Password, &user.EmpId)
	if err != nil {
		logrus.WithError(err).Warn("unable to Select doc")
		return errors.New("unable to insert doc metadata"), nil
	}
	return nil, user
}
func (d *Database) FindUserRole(tx context.Context, empid int) (error, int) {
	var role int
	err := d.database.Conn.
		QueryRow(context.Background(), "select role from temployeedetails where empid=$1", empid).Scan(&role)
	if err != nil {
		logrus.WithError(err).Warn("unable to Select doc")
		return errors.New("unable to insert doc metadata"), 0
	}
	return nil, role
}
