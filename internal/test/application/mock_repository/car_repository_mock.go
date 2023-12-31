package mock_repository

import (
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type CarRepositoryMock struct {
	mock.Mock
}

func (r *CarRepositoryMock) FindOwnerCar(ownerId string) (user.IUser, error) {
	args := r.Called(ownerId)
	return args.Get(0).(user.IUser), args.Error(1)
}

func (r *CarRepositoryMock) Register(iUser user.IUser) (user.IUser, error) {
	args := r.Called(iUser)
	return args.Get(0).(user.IUser), args.Error(1)
}
