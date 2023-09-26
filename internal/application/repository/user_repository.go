package repository

import "github.com/danyukod/decarona-register/internal/domain/user"

type UserRepositoryInterface interface {
	Save(user user.IUser) (user.IUser, error)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user user.IUser) (user.IUser, error) {

	return nil, nil
}
