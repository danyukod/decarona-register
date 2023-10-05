package repository

import "github.com/danyukod/decarona-register/internal/domain/user"

type UserRepositoryInterface interface {
	FindById(id string) (user.IUser, error)
	Update(user user.IUser) (user.IUser, error)
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
