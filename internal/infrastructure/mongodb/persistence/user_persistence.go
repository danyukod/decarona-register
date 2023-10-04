package persistence

import (
	"context"
	"github.com/danyukod/decarona-register/internal/domain/car"
	document_domain "github.com/danyukod/decarona-register/internal/domain/document"
	"github.com/danyukod/decarona-register/internal/domain/user"
	"github.com/danyukod/decarona-register/internal/infrastructure/mongodb/persistence/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "user"

type UserPersistenceInterface interface {
	Save(user user.IUser) (string, error)
	FindByEmail(email string) (user.IUser, error)
	FindByID(id string) (user.IUser, error)
	Update(user user.IUser) (string, error)
}

type UserPersistence struct {
	*mongo.Database
}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{}
}

func (r *UserPersistence) Save(ctx context.Context, user user.IUser) (*string, error) {
	userDocument := document.FromUser(user)
	id, err := r.Database.Collection(CollectionName).InsertOne(ctx, userDocument)
	if err != nil {
		return nil, err
	}
	idString := getHexObjectIdString(id.InsertedID)
	return &idString, nil
}

func (r *UserPersistence) FindByEmail(ctx context.Context, email string) (user.IUser, error) {
	var userDocument document.UserDocument
	err := r.Database.Collection(CollectionName).FindOne(ctx, document.UserDocument{Email: email}).Decode(&userDocument)
	if err != nil {
		return nil, err
	}

	return toDomainUser(userDocument)
}

func getHexObjectIdString(id interface{}) string {
	return id.(primitive.ObjectID).Hex()
}

func toDomainUser(userDocument document.UserDocument) (user.IUser, error) {
	carDomainList, err := toCarDomainList(userDocument.GetCars())
	if err != nil {
		return nil, err
	}
	documentDomainList, err := toDocumentDomainList(userDocument.GetDocuments())
	if err != nil {
		return nil, err
	}
	userDomain, err := user.NewUser(userDocument.GetName(), userDocument.GetEmail(), userDocument.GetGender(), userDocument.GetPassword(), documentDomainList)
	for _, carDomain := range carDomainList {
		userDomain.AddCar(carDomain)
	}
	if err != nil {
		return nil, err
	}
	userDomain.SetID(userDocument.GetID())
	return userDomain, nil
}

func toCarDomainList(carDocument []document.ICarDocument) ([]car.ICar, error) {
	var carDomainList []car.ICar
	for _, doc := range carDocument {
		carDomain, err := car.NewCar(doc.GetModel(), doc.GetBrand(), doc.GetColor(), doc.GetPlate(), doc.GetDescription(), doc.GetYear())
		if err != nil {
			return nil, err
		}
		carDomainList = append(carDomainList, carDomain)
	}
	return carDomainList, nil
}

func toDocumentDomainList(userDocDocuments []document.IUserDocDocument) ([]document_domain.IDocument, error) {
	var documentDomainList []document_domain.IDocument
	for _, doc := range userDocDocuments {
		documentDomain, err := document_domain.FromText(doc.GetDocType(), doc.GetDocNumber())
		if err != nil {
			return nil, err
		}
		documentDomainList = append(documentDomainList, documentDomain)
	}
	return documentDomainList, nil
}
