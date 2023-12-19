package repository

import (
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/danyukod/decarona-register/internal/infrastructure/mongodb/persistence"
)

type CarRepositoryInterface interface {
	FindOwnerCar(ownerId string) (user.IUser, error)
	Register(userCars user.IUser) (user.IUser, error)
}

type CarRepository struct {
	persistence.UserPersistenceInterface
}

func NewCarRepository(persistence persistence.UserPersistenceInterface) *CarRepository {
	return &CarRepository{
		persistence,
	}
}

func (r *CarRepository) FindOwnerCar(ownerId string) (user.IUser, error) {
	return nil, nil
}

func (r *CarRepository) Register(car car.ICar) (car.ICar, error) {
	return nil, nil
}
