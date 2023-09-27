package document

import (
	"errors"
	"strings"
)

const (
	CPF = "CPF"
	CNH = "CNH"
)

type IDocument interface {
	GetType() string
	GetNumber() string
	Validate() error
}

func FromText(documentType string, documentNumber string) (IDocument, error) {
	switch strings.ToUpper(documentType) {
	case CPF:
		return NewCPFDocument(documentNumber)
	case CNH:
		return NewCNHDocument(documentNumber)
	default:
		return nil, errors.New("invalid document type")
	}
}
