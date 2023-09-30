package commands

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r *UserRepositoryMock) Register(iUser user.IUser) (user.IUser, error) {
	args := r.Called(iUser)
	return args.Get(0).(user.IUser), args.Error(1)
}

func TestNewRegisterUserService(t *testing.T) {
	repository := UserRepositoryMock{}
	service := NewRegisterUserService(&repository)
	assert.NotNil(t, service)

	t.Run("should return a user with success", func(t *testing.T) {

		registerUserDto := mockNewUserDto()

		returnedUserDomain := mockNewUserDomain(registerUserDto)
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

func mockNewUserDto() dto.RegisterUserDTO {
	return dto.RegisterUserDTO{
		Name:     "John Doe",
		Email:    "jd@email.com",
		Gender:   "M",
		Password: "12345678",
		Documents: []dto.DocumentDTO{
			{
				Number:  "66329748055",
				DocType: "CPF",
			},
			{
				Number:  "90496138465",
				DocType: "CNH",
			},
		},
	}
}

func mockNewUserDomain(dto dto.RegisterUserDTO) user.IUser {
	documents := mockNewDocuments(dto.Documents)
	userDomain, _ := user.NewUser(dto.Name, dto.Email, dto.Gender, dto.Password, documents)
	return userDomain
}

func mockNewDocuments(documents []dto.DocumentDTO) []document.IDocument {
	var documentsDomain []document.IDocument
	for _, doc := range documents {
		documentDomain, _ := document.FromText(doc.DocType, doc.Number)
		documentsDomain = append(documentsDomain, documentDomain)
	}
	return documentsDomain
}
