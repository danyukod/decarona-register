package car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar(t *testing.T) {
	model := "Fiesta"
	brand := "Ford"
	color := "Red"
	description := "A cool car"
	document := "ABC1234"
	year := 2020

	t.Run("should create new car successfuly", func(t *testing.T) {
		car, err := NewCar(model, brand, color, document, &description, year)

		assert.Nil(t, err)
		assert.NotNil(t, car)
		assert.Equal(t, model, car.GetModel())
		assert.Equal(t, brand, car.GetBrand())
		assert.Equal(t, color, car.GetColor())
		assert.Equal(t, year, car.GetYear())
		assert.Equal(t, description, *car.GetDescription())
	})

	t.Run("should return error when model is empty", func(t *testing.T) {
		car, err := NewCar("", brand, color, document, &description, year)

		assert.NotNil(t, err)
		assert.Nil(t, car)
		assert.Equal(t, "invalid model", err.Error())
	})

	t.Run("should return error when brand is empty", func(t *testing.T) {
		car, err := NewCar(model, "", color, document, &description, year)

		assert.NotNil(t, err)
		assert.Nil(t, car)
		assert.Equal(t, "invalid brand", err.Error())
	})

	t.Run("should return error when color is empty", func(t *testing.T) {
		car, err := NewCar(model, brand, "", document, &description, year)

		assert.NotNil(t, err)
		assert.Nil(t, car)
		assert.Equal(t, "invalid color", err.Error())
	})

	t.Run("should return error when year is empty", func(t *testing.T) {
		car, err := NewCar(model, brand, color, document, &description, 0)

		assert.NotNil(t, err)
		assert.Nil(t, car)
		assert.Equal(t, "invalid year", err.Error())
	})

	t.Run("should return error when plate is empty", func(t *testing.T) {
		car, err := NewCar(model, brand, color, "", &description, year)

		assert.NotNil(t, err)
		assert.Nil(t, car)
		assert.Equal(t, "invalid plate", err.Error())
	})
}
