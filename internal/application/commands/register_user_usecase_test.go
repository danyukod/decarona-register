package commands

import (
	"github.com/danyukod/decarona-register/internal/test/application/mock_commands/mock_dto"
	"github.com/danyukod/decarona-register/internal/test/application/mock_repository"
	"github.com/danyukod/decarona-register/internal/test/domain/mock_user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewRegisterUserService(t *testing.T) {
	repository := mock_repository.UserRepositoryMock{}
	service := NewRegisterUserService(&repository)
	assert.NotNil(t, service)

	t.Run("should return a mock_user with success", func(t *testing.T) {

		registerUserDto := mock_dto.MockNewUserDto()

		returnedUserDomain := mock_user.MockNewUserDomain(registerUserDto)
		returnedUserDomain.SetID("1")

		repository.On("Register", mock.Anything).Return(returnedUserDomain, nil).Once()

		userDomain, err := service.Execute(registerUserDto)

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
		assert.Equal(t, "1", userDomain.GetID())
		assert.Equal(t, registerUserDto.Name, userDomain.GetName())
		assert.Equal(t, registerUserDto.Email, userDomain.GetEmail())
		assert.Equal(t, registerUserDto.Gender, userDomain.GetGender())
		assert.True(t, returnedUserDomain.ValidatePassword(registerUserDto.Password))
		assert.Equal(t, len(registerUserDto.Documents), len(userDomain.GetDocuments()))
		assert.Equal(t, registerUserDto.Documents[0].Number, userDomain.GetDocuments()[0].GetNumber())
		assert.Equal(t, registerUserDto.Documents[0].DocType, userDomain.GetDocuments()[0].GetType())
		assert.Equal(t, registerUserDto.Documents[1].Number, userDomain.GetDocuments()[1].GetNumber())
		assert.Equal(t, registerUserDto.Documents[1].DocType, userDomain.GetDocuments()[1].GetType())
	})

}
