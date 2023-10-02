package mock_car

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/domain/car"
)

func MockNewCarDomain(dto dto.RegisterCarDTO) car.ICar {
	carDomain, _ := car.NewCar(dto.Model, dto.Brand, dto.Color, dto.Plate, dto.Description, dto.Year)
	return carDomain
}
