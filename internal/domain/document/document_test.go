package document

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocument(t *testing.T) {
	t.Run("should create new cpf document successfuly", func(t *testing.T) {
		documentType := "CPF"
		documentNumber := "66329748055"

		document, err := FromText(documentType, documentNumber)

		assert.Nil(t, err)
		assert.NotNil(t, document)
		assert.Equal(t, documentType, document.GetType())
		assert.Equal(t, documentNumber, document.GetNumber())
	})

	t.Run("should create new cnh document successfuly", func(t *testing.T) {
		documentType := "CNH"
		documentNumber := "90496138465"

		document, err := FromText(documentType, documentNumber)

		assert.Nil(t, err)
		assert.NotNil(t, document)
		assert.Equal(t, documentType, document.GetType())
		assert.Equal(t, documentNumber, document.GetNumber())
	})

	t.Run("should return error when document type is invalid", func(t *testing.T) {
		documentType := "INVALID"
		documentNumber := "66329748055"

		document, err := FromText(documentType, documentNumber)

		assert.NotNil(t, err)
		assert.Nil(t, document)
		assert.Equal(t, "invalid document type", err.Error())
	})

	t.Run("should return error when cpf number is invalid", func(t *testing.T) {
		documentType := "CPF"
		documentNumber := "123456789"

		document, err := FromText(documentType, documentNumber)

		assert.NotNil(t, err)
		assert.Nil(t, document)
		assert.Equal(t, "invalid cpf", err.Error())
	})

	t.Run("should return error when cnh number is invalid", func(t *testing.T) {
		documentType := "CNH"
		documentNumber := "123456789012"

		document, err := FromText(documentType, documentNumber)

		assert.NotNil(t, err)
		assert.Nil(t, document)
		assert.Equal(t, "invalid cnh", err.Error())
	})

}
