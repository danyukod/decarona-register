package commands

import (
	"github.com/danyukod/decarona-register/internal/test/application/mock_commands/mock_dto"
	"github.com/danyukod/decarona-register/internal/test/application/mock_repository"
	"github.com/danyukod/decarona-register/internal/test/domain/mock_car"
	"github.com/danyukod/decarona-register/internal/test/domain/mock_user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewRegisterCarService(t *testing.T) {
	repository := mock_repository.CarRepositoryMock{}
	service := NewRegisterCarService(&repository)
	assert.NotNil(t, service)

	t.Run("should return a owner with registered mock_car with success", func(t *testing.T) {

		ownerCarDomain := mock_user.MockNewUserDomain(mock_dto.MockNewUserDto())
		ownerCarDomain.SetID("1")

		registerCarDto := mock_dto.MockNewCarDto()

		repository.On("FindOwnerCar", registerCarDto.OwnerId).Return(ownerCarDomain, nil).Once()

		carDomain := mock_car.MockNewCarDomain(registerCarDto)
		ownerCarDomainWithCar := mock_user.MockNewUserDomain(mock_dto.MockNewUserDto())
		ownerCarDomainWithCar.SetID("1")
		ownerCarDomainWithCar.AddCar(carDomain)

		repository.On("Register", mock.Anything).Return(ownerCarDomainWithCar, nil).Once()

		userDomain, err := service.Execute(registerCarDto)

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
		assert.Equal(t, "1", userDomain.GetID())
		assert.Equal(t, registerCarDto.OwnerId, userDomain.GetID())
		assert.Equal(t, registerCarDto.Model, userDomain.GetCars()[0].GetModel())
		assert.Equal(t, registerCarDto.Brand, userDomain.GetCars()[0].GetBrand())
		assert.Equal(t, registerCarDto.Year, userDomain.GetCars()[0].GetYear())
		assert.Equal(t, registerCarDto.Color, userDomain.GetCars()[0].GetColor())
		assert.Equal(t, registerCarDto.Plate, userDomain.GetCars()[0].GetPlate())
		assert.Equal(t, *registerCarDto.Description, *userDomain.GetCars()[0].GetDescription())
	})
}
