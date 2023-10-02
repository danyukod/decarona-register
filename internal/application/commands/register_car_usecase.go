package commands

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/application/repository"
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/user"
)

type RegisterCarUsecase interface {
	Execute(dto dto.RegisterCarDTO) (user.IUser, error)
}

type RegisterCarService struct {
	repository.CarRepositoryInterface
}

func NewRegisterCarService(repository repository.CarRepositoryInterface) *RegisterCarService {
	return &RegisterCarService{
		repository,
	}
}

func (r *RegisterCarService) Execute(dto dto.RegisterCarDTO) (user.IUser, error) {
	userDomain, err := r.FindOwnerCar(dto.OwnerId)
	if err != nil {
		return nil, err
	}

	carDomain, err := car.NewCar(dto.Model, dto.Brand, dto.Color, dto.Plate, dto.Description, dto.Year)
	if err != nil {
		return nil, err
	}

	userDomain.AddCar(carDomain)

	userDomain, err = r.Register(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomain, nil
}
