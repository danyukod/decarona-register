package user

import (
	"errors"
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/document"
	"golang.org/x/crypto/bcrypt"
)

const (
	male   = "M"
	female = "F"
)

type IUser interface {
	GetID() string
	SetID(id string)
	GetName() string
	GetEmail() string
	GetPassword() string
	GetGender() string
	GetDocuments() []document.IDocument
	AddCar(car car.ICar)
	Validate() error
	ValidatePassword(password string) bool
}

type user struct {
	id        string
	name      string
	email     string
	gender    string
	documents []document.IDocument
	cars      []car.ICar
	password  string
}

func NewUser(name, email, gender, password string, documents []document.IDocument) (IUser, error) {
	userDomain := user{
		name:      name,
		email:     email,
		documents: documents,
		gender:    gender,
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

func (u *user) GetGender() string {
	return u.gender
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
	if !u.validateGender() {
		return errors.New("undefined gender")
	}
	return nil
}

func (u *user) validateGender() bool {
	switch u.gender {
	case male:
		return true
	case female:
		return true
	default:
		return false
	}
}

func (u *user) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u *user) AddCar(car car.ICar) {
	u.cars = append(u.cars, car)
}
