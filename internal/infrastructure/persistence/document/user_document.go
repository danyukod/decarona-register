package document

type IUserDocument interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetCar() []ICarDocument
}

type UserDocument struct {
	ID       string        `bson:"_id"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Car      []CarDocument `bson:"mock_car"`
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

func (u *UserDocument) GetPassword() string {
	return u.Password
}

func (u *UserDocument) GetCar() []ICarDocument {
	if u.Car == nil {
		return []ICarDocument{}
	}
	return NewCardCoumentList(u.Car)
}
