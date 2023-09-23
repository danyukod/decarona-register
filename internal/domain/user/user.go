package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	GetID() string
	SetID(id string)
	GetName() string
	GetEmail() string
	GetDocument() string
	GetPassword() string
	Validate() error
	ValidatePassword(password string) bool
}

type user struct {
	id       string
	name     string
	email    string
	document string
	password string
}

func NewUser(name, email, document, password string) (IUser, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userDomain := user{
		name:     name,
		email:    email,
		document: document,
		password: string(hash),
	}
	if err := userDomain.Validate(); err != nil {
		return nil, err
	}
	return &userDomain, nil
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

func (u *user) GetDocument() string {
	return u.document
}

func (u *user) Validate() error {
	if u.name == "" {
		return errors.New("invalid name")
	}
	if u.email == "" {
		return errors.New("invalid email")
	}
	if u.document == "" {
		return errors.New("invalid document")
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
