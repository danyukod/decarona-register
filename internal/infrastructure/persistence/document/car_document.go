package document

type ICarDocument interface {
	GetBrand() string
	GetModel() string
	GetColor() string
	GetDocument() string
	GetYear() int
}

type CarDocument struct {
	Model       string  `bson:"model"`
	Brand       string  `bson:"brand"`
	Year        int     `bson:"year"`
	Color       string  `bson:"color"`
	Document    string  `bson:"document"`
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

func (c *CarDocument) GetDocument() string {
	return c.Document
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
