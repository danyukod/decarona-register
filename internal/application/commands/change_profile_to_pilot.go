package commands

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/application/repository"
	"github.com/danyukod/decarona-register/internal/domain/document"
)

type ChangeProfileToPilotUsecase interface {
	Execute(dto dto.ChangeProfileToPilotDTO) error
}

type ChangeProfileToPilotService struct {
	repository.UserRepositoryInterface
}

func NewChangeProfileToPilotService(repository repository.UserRepositoryInterface) *ChangeProfileToPilotService {
	return &ChangeProfileToPilotService{
		repository,
	}
}

func (r *ChangeProfileToPilotService) Execute(dto dto.ChangeProfileToPilotDTO) error {
	userDomain, err := r.FindById(dto.UserId)
	if err != nil {
		return err
	}

	cnh, err := document.FromText(document.CNH, dto.Cnh)
	if err != nil {
		return err
	}

	userDomain.AddDocument(cnh)

	_, err = r.Update(userDomain)
	if err != nil {
		return err
	}

	return nil
}
