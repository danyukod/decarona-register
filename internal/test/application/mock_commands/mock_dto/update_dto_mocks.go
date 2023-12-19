package mock_dto

import "github.com/danyukod/decarona-register/internal/application/commands/dto"

func MockNewChangeProfileToPilotDTO() dto.ChangeProfileToPilotDTO {
	return dto.ChangeProfileToPilotDTO{
		UserId: "1",
		Cnh:    "90496138465",
	}
}
