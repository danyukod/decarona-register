package repository

import (
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/user"
)

type CarRepositoryInterface interface {
	FindOwnerCar(ownerId string) (user.IUser, error)
	Register(userCars user.IUser) (user.IUser, error)
}

type CarRepository struct {
}

func NewCarRepository() *CarRepository {
	return &CarRepository{}
}

func (r *CarRepository) FindOwnerCar(ownerId string) (user.IUser, error) {
	return nil, nil
}

func (r *CarRepository) Register(car car.ICar) (car.ICar, error) {
	return nil, nil
}
