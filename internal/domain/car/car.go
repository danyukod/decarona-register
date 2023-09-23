package car

import "errors"

type ICar interface {
	GetID() string
	SetID(id string)
	GetModel() string
	GetBrand() string
	GetYear() int
	GetColor() string
	GetDescription() *string
}

type car struct {
	id          string
	model       string
	brand       string
	year        int
	color       string
	description *string
}

func NewCar(model, brand, color string, description *string, year int) (ICar, error) {
	carDomain := car{
		model:       model,
		brand:       brand,
		year:        year,
		color:       color,
		description: description,
	}
	if err := carDomain.Validate(); err != nil {
		return nil, err
	}
	return &carDomain, nil
}

func (c *car) GetID() string {
	return c.id
}

func (c *car) SetID(id string) {
	c.id = id
}

func (c *car) GetModel() string {
	return c.model
}

func (c *car) GetBrand() string {
	return c.brand
}

func (c *car) GetYear() int {
	return c.year
}

func (c *car) GetColor() string {
	return c.color
}

func (c *car) GetDescription() *string {
	return c.description
}

func (c *car) Validate() error {
	if c.model == "" {
		return errors.New("invalid model")
	}
	if c.brand == "" {
		return errors.New("invalid brand")
	}
	if c.year == 0 {
		return errors.New("invalid year")
	}
	if c.color == "" {
		return errors.New("invalid color")
	}
	return nil
}
