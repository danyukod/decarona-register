package repository

import "github.com/danyukod/decarona-register/internal/domain/user"

type UserRepositoryInterface interface {
	Register(user user.IUser) (user.IUser, error)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Register(user user.IUser) (user.IUser, error) {

	return nil, nil
}
