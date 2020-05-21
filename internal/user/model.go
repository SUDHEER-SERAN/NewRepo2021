package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}
