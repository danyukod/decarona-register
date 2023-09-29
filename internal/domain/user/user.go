package user

import (
	"errors"
	"github.com/danyukod/decarona-register/internal/domain/document"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	GetID() string
	SetID(id string)
	GetName() string
	GetEmail() string
	GetDocuments() []document.IDocument
	GetPassword() string
	Validate() error
	ValidatePassword(password string) bool
}

type user struct {
	id        string
	name      string
	email     string
	documents []document.IDocument
	password  string
}

func NewUser(name, email, password string, documents []document.IDocument) (IUser, error) {
	userDomain := user{
		name:      name,
		email:     email,
		documents: documents,
		password:  encryptPassword(password),
	}
	if err := userDomain.Validate(); err != nil {
		return nil, err
	}
	return &userDomain, nil
}

func encryptPassword(password string) string {
	if password == "" {
		return password
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password
	}
	return string(hash)
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) GetID() string {
	return u.id
}

func (u *user) SetID(id string) {
	u.id = id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetDocuments() []document.IDocument {
	return u.documents
}

func (u *user) Validate() error {
	if u.name == "" {
		return errors.New("invalid name")
	}
	if u.email == "" {
		return errors.New("invalid email")
	}
	if u.password == "" {
		return errors.New("invalid password")
	}
	return nil
}

func (u *user) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}
