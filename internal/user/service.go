package user

type Service interface {
	CreateUser()
}
type Repository interface {
	CreateUserDB() string
}

type userService struct {
	repo Repository
}

func NewUserService(r Repository) Service {
	s := &userService{
		repo: r,
	}
	return s
}

func (s *userService) CreateUser() {
	s.repo.CreateUserDB()
}
