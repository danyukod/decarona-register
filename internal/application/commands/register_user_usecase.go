package commands

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/application/repository"
	"github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/danyukod/decarona-register/internal/domain/user"
)

type RegisterUserUsecase interface {
	Execute(dto dto.RegisterUserDTO) (user.IUser, error)
}

type RegisterUserService struct {
	repository.UserRepositoryInterface
}

func NewRegisterUserService(repository repository.UserRepositoryInterface) *RegisterUserService {
	return &RegisterUserService{
		repository,
	}
}

func (r *RegisterUserService) Execute(dto dto.RegisterUserDTO) (user.IUser, error) {
	var documents []document.IDocument

	for _, documentDTO := range dto.Documents {
		docFromText, err := document.FromText(documentDTO.DocType, documentDTO.Number)
		if err != nil {
			return nil, err
		}
		err = docFromText.Validate()
		if err != nil {
			return nil, err
		}
		documents = append(documents, docFromText)
	}

	userDomain, err := user.NewUser(dto.Name, dto.Email, dto.Password, documents)
	if err != nil {
		return nil, err
	}

	domain, err := r.Register(userDomain)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
