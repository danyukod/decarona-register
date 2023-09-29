package user

import (
	"github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {

	documentNumber1 := "12345678900"
	documentNumber2 := "12345678901"
	documentType1 := "CPF"
	documentType2 := "CNH"

	name := "John Doe"
	email := "jd@email.com"
	password := "123456"

	doc1, err := document.FromText(documentType1, documentNumber1)
	assert.Nil(t, err)
	doc2, err := document.FromText(documentType2, documentNumber2)
	assert.Nil(t, err)

	var documentList []document.IDocument

	documentList = append(documentList, doc1)
	documentList = append(documentList, doc2)

	t.Run("should create new user successfuly", func(t *testing.T) {

		user, err := NewUser(name, email, password, documentList)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.GetName())
		assert.Equal(t, email, user.GetEmail())
		assert.Equal(t, documentList, user.GetDocuments())
		assert.NotEqual(t, password, user.GetPassword())
		assert.True(t, user.ValidatePassword(password))
	})

	t.Run("should return error when name is empty", func(t *testing.T) {

		user, err := NewUser("", email, password, documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid name", err.Error())
	})

	t.Run("should return error when email is empty", func(t *testing.T) {

		user, err := NewUser(name, "", password, documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("should return error when password is empty", func(t *testing.T) {

		user, err := NewUser(name, email, "", documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid password", err.Error())
	})

}
