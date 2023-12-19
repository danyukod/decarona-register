package document

import (
	"errors"
	"github.com/danyukod/decarona-register/pkg"
)

func NewCPFDocument(number string) (IDocument, error) {
	cpf := cpfDocument{
		number: number,
	}
	if err := cpf.Validate(); err != nil {
		return nil, err
	}
	return cpf, nil
}

type cpfDocument struct {
	number string
}

func (c cpfDocument) GetType() string {
	return CPF
}

func (c cpfDocument) GetNumber() string {
	return c.number
}

func (c cpfDocument) Validate() error {
	if pkg.IsCPF(c.number) {
		return nil
	}
	return errors.New("invalid cpf")
}
