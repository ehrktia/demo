package service

import (
	"context"

	"github.com/ehrktia/demo/entity"
)

//go:generate mockery --name UserRepository --with-expecter=true --outpkg mocks --output ../mocks
type UserRepository interface {
	GetUserById(ctx context.Context, id int) (entity.User, error)
}

type UserService struct {
	UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		ur,
	}
}

func (us *UserService) GetUserById(ctx context.Context, id int) (entity.User, error) {
	return us.UserRepository.GetUserById(ctx, id)
}
