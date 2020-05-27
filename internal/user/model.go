package user

import (
	"time"
)

type User struct {
	EmpId    int    `json:"empId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmployeeDetails struct {
	EmpId     int    `json:"empId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	MobileNo  int    `json:"mobileNo"`
	Address   string `json:"address"`
	Role      int    `json:"role"`
}

type Token struct {
	Type        string    `json:"type"`
	AccessToken string    `json:"token"`
	Expires     time.Time `json:"expires"`
}
