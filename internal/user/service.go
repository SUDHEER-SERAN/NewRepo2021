package user

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, userDetails User) error
	AuthenticateUser(ctx context.Context, userDetails User) error
}

type userService struct {
	repo *Database
}

func NewUserService(r *Database) Service {
	s := &userService{
		repo: r,
	}
	return s
}

func (s *userService) CreateUser(ctx context.Context, userDetails User) error {
	ePwd, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.DefaultCost)

	if err != nil {
		logrus.WithError(err).Error("unable to encrypt password")
		return errors.New("unable to encrypt password")
	}
	userDetails.Password = string(ePwd)
	if err := s.repo.CreateUserRepo(ctx, userDetails); err != nil {
		logrus.WithError(err).Error("unable to create user")
		return errors.New("unable to create user")
	}
	return nil
}

func (s *userService) AuthenticateUser(ctx context.Context, userDetails User) error {
	return nil
}
