package persistence

import "github.com/danyukod/decarona-register/internal/domain/user"

type UserPersistenceInterface interface {
	Save(user user.IUser) (user.IUser, error)
}

type UserPersistence struct {
}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{}
}

func (r *UserPersistence) Save(user user.IUser) (user.IUser, error) {
	return nil, nil
}
