package document

import (
	"github.com/danyukod/decarona-register/internal/domain/car"
	"github.com/danyukod/decarona-register/internal/domain/user"
)

type IUserDocument interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetGender() string
	GetPassword() string
	GetDocuments() []IUserDocDocument
	GetCars() []ICarDocument
}

type UserDocument struct {
	ID        string             `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Gender    string             `bson:"gender"`
	Password  string             `bson:"password"`
	Documents []IUserDocDocument `bson:"documents"`
	Cars      []CarDocument      `bson:"mock_car"`
}

func (u *UserDocument) GetID() string {
	return u.ID
}

func (u *UserDocument) GetName() string {
	return u.Name
}

func (u *UserDocument) GetEmail() string {
	return u.Email
}

func (u *UserDocument) GetGender() string {
	return u.Gender
}

func (u *UserDocument) GetDocuments() []IUserDocDocument {
	return u.Documents
}

func (u *UserDocument) GetPassword() string {
	return u.Password
}

func (u *UserDocument) GetCars() []ICarDocument {
	if u.Cars == nil {
		return []ICarDocument{}
	}
	return NewCardCoumentList(u.Cars)
}

func FromUser(user user.IUser) UserDocument {
	return UserDocument{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Cars:     carDocumentListFromCars(user.GetCars()),
	}
}

func carDocumentListFromCars(cars []car.ICar) []CarDocument {
	var carDocumentList []CarDocument
	for _, carDomain := range cars {
		carDocumentList = append(carDocumentList, documentFromCar(carDomain))
	}
	return carDocumentList
}
