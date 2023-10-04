package document

import "github.com/danyukod/decarona-register/internal/domain/car"

type ICarDocument interface {
	GetBrand() string
	GetModel() string
	GetColor() string
	GetYear() int
	GetDescription() *string
	GetPlate() string
}

type CarDocument struct {
	Model       string  `bson:"model"`
	Brand       string  `bson:"brand"`
	Year        int     `bson:"year"`
	Color       string  `bson:"color"`
	Plate       string  `bson:"document"`
	Description *string `bson:"description"`
}

func (c *CarDocument) GetModel() string {
	return c.Model
}

func (c *CarDocument) GetBrand() string {
	return c.Brand
}

func (c *CarDocument) GetYear() int {
	return c.Year
}

func (c *CarDocument) GetColor() string {
	return c.Color
}

func (c *CarDocument) GetPlate() string {
	return c.Plate
}

func (c *CarDocument) GetDescription() *string {
	return c.Description
}

func NewCardCoumentList(car []CarDocument) []ICarDocument {
	var carDocumentList []ICarDocument
	for _, car := range car {
		carDocumentList = append(carDocumentList, &car)
	}
	return carDocumentList
}

func documentFromCar(iCar car.ICar) CarDocument {
	return CarDocument{
		Model:       iCar.GetModel(),
		Brand:       iCar.GetBrand(),
		Year:        iCar.GetYear(),
		Color:       iCar.GetColor(),
		Plate:       iCar.GetPlate(),
		Description: iCar.GetDescription(),
	}
}
