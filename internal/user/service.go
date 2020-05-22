package user

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	key = "secretKey"
)

type Service interface {
	CreateUser(ctx context.Context, userDetails User) error
	AuthenticateUser(ctx context.Context, userDetails User) (string, error)
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

func (s *userService) AuthenticateUser(ctx context.Context, userDetails User) (string, error) {

	err, rec := s.repo.FindUser(ctx, userDetails)
	if err != nil {
		logrus.WithError(err).Error("Unable to fetch the record")
		return "", errors.New("Cant be logined")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(rec.Password), []byte(userDetails.Password)); err != nil {
		logrus.WithError(err).Error("Canot Access")
		return "", errors.New("unable to access the Apllication")
	}
	tockenExpiration := time.Now().Add(time.Hour * 72)

	tokenString, err := s.getToken(userDetails, tockenExpiration)
	if err != nil {
		logrus.WithError(err).Error("Canot generate the Jwt token")
		return "", errors.New("unable to generate the JWT Token")
	}
	return tokenString, nil
}

func (s *userService) getToken(user User, expiration time.Time) (string, error) {
	// CreateEntry the token
	token := jwt.New(jwt.SigningMethodHS256)

	// CreateEntry a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = expiration.Unix()

	// Sign the token with our secret
	return token.SignedString([]byte(key))
}
