package user

import (
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {

	documentNumber1 := "66329748055"
	documentNumber2 := "90496138465"
	documentType1 := "CPF"
	documentType2 := "CNH"

	name := "John Doe"
	email := "jd@email.com"
	password := "123456"
	gender := "M"

	doc1, err := document.FromText(documentType1, documentNumber1)
	assert.Nil(t, err)
	doc2, err := document.FromText(documentType2, documentNumber2)
	assert.Nil(t, err)

	var documentList []document.IDocument

	documentList = append(documentList, doc1)
	documentList = append(documentList, doc2)

	t.Run("should create new mock_user successfuly", func(t *testing.T) {

		user, err := NewUser(name, email, gender, password, documentList)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.GetName())
		assert.Equal(t, email, user.GetEmail())
		assert.Equal(t, documentList, user.GetDocuments())
		assert.NotEqual(t, password, user.GetPassword())
		assert.True(t, user.ValidatePassword(password))
	})

	t.Run("should return error when name is empty", func(t *testing.T) {

		user, err := NewUser("", email, gender, password, documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid name", err.Error())
	})

	t.Run("should return error when email is empty", func(t *testing.T) {

		user, err := NewUser(name, "", gender, password, documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("should return error when gender is empty", func(t *testing.T) {

		user, err := NewUser(name, email, "", password, documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "undefined gender", err.Error())
	})

	t.Run("should return error when password is empty", func(t *testing.T) {

		user, err := NewUser(name, email, gender, "", documentList)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid password", err.Error())
	})

	t.Run("should add new car to user", func(t *testing.T) {

		user, err := NewUser(name, email, gender, password, documentList)
		assert.Nil(t, err)
		assert.NotNil(t, user)

		newCar, err := car.NewCar("model", "brand", "color", "plate", nil, 2020)
		assert.Nil(t, err)
		assert.NotNil(t, newCar)

		user.AddCar(newCar)

		assert.Equal(t, 1, len(user.GetCars()))
	})

	t.Run("should add new document to user", func(t *testing.T) {

		doc, err := document.FromText("CPF", "66329748055")
		assert.Nil(t, err)
		assert.NotNil(t, doc)
		var documents []document.IDocument
		documents = append(documents, doc)
		user, err := NewUser(name, email, gender, password, documents)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		doc, err = document.FromText("CNH", "90496138465")
		assert.Nil(t, err)
		assert.NotNil(t, doc)

		user.AddDocument(doc)

		assert.Equal(t, 2, len(user.GetDocuments()))

	})

}
