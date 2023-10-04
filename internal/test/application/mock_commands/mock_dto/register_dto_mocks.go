package mock_dto

import (
	"github.com/danyukod/decarona-register/internal/application/commands/dto"
	"github.com/danyukod/decarona-register/internal/domain/document"
)

func MockNewUserDto() dto.RegisterUserDTO {
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

func MockNewCarDto() dto.RegisterCarDTO {
	description := "This is a description"
	return dto.RegisterCarDTO{
		OwnerId:     "1",
		Model:       "Fiesta",
		Brand:       "Ford",
		Year:        2020,
		Color:       "Red",
		Plate:       "ABC1234",
		Description: &description,
	}
}

func MockNewDocuments(documents []dto.DocumentDTO) []document.IDocument {
	var documentsDomain []document.IDocument
	for _, doc := range documents {
		documentDomain, _ := document.FromText(doc.DocType, doc.Number)
		documentsDomain = append(documentsDomain, documentDomain)
	}
	return documentsDomain
}

func MockNewUserWithoutCnhDto() dto.RegisterUserDTO {
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
		},
	}
}
