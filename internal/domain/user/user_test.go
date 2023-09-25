package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {

	name := "John Doe"
	email := "jd@email.com"
	document := "12345678900"
	password := "123456"

	t.Run("should create new user successfuly", func(t *testing.T) {

		user, err := NewUser(name, email, document, password)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.GetName())
		assert.Equal(t, email, user.GetEmail())
		assert.Equal(t, document, user.GetDocument())
		assert.NotEqual(t, password, user.GetPassword())
		assert.True(t, user.ValidatePassword(password))
	})

	t.Run("should return error when name is empty", func(t *testing.T) {

		user, err := NewUser("", email, document, password)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid name", err.Error())
	})

	t.Run("should return error when email is empty", func(t *testing.T) {

		user, err := NewUser(name, "", document, password)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("should return error when document is empty", func(t *testing.T) {

		user, err := NewUser(name, email, "", password)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid document", err.Error())
	})

	t.Run("should return error when password is empty", func(t *testing.T) {
		user, err := NewUser(name, email, document, "")

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid password", err.Error())
	})

}
