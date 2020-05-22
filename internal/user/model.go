package user

import (
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

type Token struct {
	Type        string    `json:"type"`
	AccessToken string    `json:"token"`
	Expires     time.Time `json:"expires"`
}
