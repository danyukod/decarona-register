package repository

import (
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/danyukod/decarona-register/internal/infrastructure/mongodb/persistence"
)

type UserRepositoryInterface interface {
	FindById(id string) (user.IUser, error)
	Update(user user.IUser) (user.IUser, error)
	Register(user user.IUser) (user.IUser, error)
}

type UserRepository struct {
	persistence.UserPersistenceInterface
}

func NewUserRepository(persistence persistence.UserPersistenceInterface) *UserRepository {
	return &UserRepository{
		persistence,
	}
}

func (r *UserRepository) Register(user user.IUser) (user.IUser, error) {

	return nil, nil
}
