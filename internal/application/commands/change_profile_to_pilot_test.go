package commands

import (
	"github.com/danyukod/decarona-register/internal/test/application/mock_commands/mock_dto"
	"github.com/danyukod/decarona-register/internal/test/application/mock_repository"
	"github.com/danyukod/decarona-register/internal/test/domain/mock_user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewChangeProfileToPilotService(t *testing.T) {
	repository := mock_repository.UserRepositoryMock{}
	service := NewChangeProfileToPilotService(&repository)
	assert.NotNil(t, service)

	t.Run("should return a changed profile user with success", func(t *testing.T) {
		userWithoutCnh := mock_dto.MockNewUserWithoutCnhDto()
		userWithoutCnhDomain := mock_user.MockNewUserDomain(userWithoutCnh)
		userWithoutCnhDomain.SetID("1")

		userWithCnh := mock_dto.MockNewUserDto()
		userWithCnhDomain := mock_user.MockNewUserDomain(userWithCnh)

		repository.On("FindById", "1").Return(userWithoutCnhDomain, nil).Once()
		repository.On("Update", mock.Anything).Return(userWithCnhDomain, nil).Once()

		pilotDTO := mock_dto.MockNewChangeProfileToPilotDTO()
		err := service.Execute(pilotDTO)
		assert.Nil(t, err)

		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "FindById", 1)
		repository.AssertNumberOfCalls(t, "Update", 1)

	})
}
