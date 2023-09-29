package document

import (
	"errors"
	"github.com/danyukod/decarona-register/pkg"
)

func NewCNHDocument(number string) (IDocument, error) {
	cnh := cnhDocument{
		number: number,
	}
	if err := cnh.Validate(); err != nil {
		return nil, err
	}
	return cnh, nil
}

type cnhDocument struct {
	number string
}

func (c cnhDocument) GetType() string {
	return CNH
}

func (c cnhDocument) GetNumber() string {
	return c.number
}

func (c cnhDocument) Validate() error {
	if pkg.IsCNH(c.number) {
		return nil
	}
	return errors.New("invalid cnh")
}
