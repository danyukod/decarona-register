package document

import "errors"

const (
	CPF = "CPF"
	CNH = "CNH"
)

type IDocument interface {
	GetType() string
	GetNumber() string
	Validate() error
}

func DocumentFromText(documentType string, documentNumber string) (IDocument, error) {
	switch documentType {
	case CPF:
		return CPFDocument{
			number: documentNumber,
		}, nil
	case CNH:
		return CNHDocument{
			number: documentNumber,
		}, nil
	default:
		return nil, errors.New("invalid document type")
	}
}

type CPFDocument struct {
	number string
}

func (c CPFDocument) GetType() string {
	return CPF
}

func (c CPFDocument) GetNumber() string {
	return c.number
}

func (c CPFDocument) Validate() error {
	return nil
}

type CNHDocument struct {
	number string
}

func (c CNHDocument) GetType() string {
	return CNH
}

func (c CNHDocument) GetNumber() string {
	return c.number
}

func (c CNHDocument) Validate() error {
	return nil
}
