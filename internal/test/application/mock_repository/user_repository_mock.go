package mock_repository

import (
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r *UserRepositoryMock) FindById(id string) (user.IUser, error) {
	args := r.Called(id)
	return args.Get(0).(user.IUser), args.Error(1)
}

func (r *UserRepositoryMock) Update(iUser user.IUser) (user.IUser, error) {
	args := r.Called(iUser)
	return args.Get(0).(user.IUser), args.Error(1)
}

func (r *UserRepositoryMock) Register(iUser user.IUser) (user.IUser, error) {
	args := r.Called(iUser)
	return args.Get(0).(user.IUser), args.Error(1)
}
