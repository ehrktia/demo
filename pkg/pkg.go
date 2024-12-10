package pkg

import "context"

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (User, error)
}

type User struct {
	Name string
	Id   int
}

type UserService struct {
	UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		ur,
	}
}

func (us *UserService) GetUserById(ctx context.Context, id int) (User, error) {

}
