package mock_user

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/danyukod/decarona-register/internal/test/application/mock_commands/mock_dto"
)

func MockNewUserDomain(dto dto.RegisterUserDTO) user.IUser {
	documents := mock_dto.MockNewDocuments(dto.Documents)
	userDomain, _ := user.NewUser(dto.Name, dto.Email, dto.Gender, dto.Password, documents)
	return userDomain
}
