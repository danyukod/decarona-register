package commands

import "github.com/danyukod/decarona-register/internal/domain/user"

type RegisterUserUsecase interface {
	Execute(dto RegisterUserDTO) (user.IUser, error)
}

type RegisterUserService struct {
}

func NewRegisterUserService() *RegisterUserService {
	return &RegisterUserService{}
}

func (r *RegisterUserService) Execute(dto RegisterUserDTO) (user.IUser, error) {
	return nil, nil
}
