package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	password := "testPassword"
	u, err := NewUser("testName", "testEmail", "testDocument", password)

	assert.Nil(t, err, "Failed to create new user: %v", err)

	assert.NotNil(t, u, "Failed to create new user: User is null")

	assert.Equal(t, "testName", u.GetName(), "Failed setattr name: Expected testName, got %v", u.GetName())

	assert.Equal(t, "testEmail", u.GetEmail(), "Failed setattr email: Expected testEmail, got %v", u.GetEmail())

	assert.Equal(t, "testDocument", u.GetDocument(), "Failed setattr document: Expected testDocument, got %v", u.GetDocument())

	assert.True(t, u.ValidatePassword(password), "Failed validate password")

}
