package generator

import (
	"faker-generator/entities"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func GenerateUsers() ([]entities.User, error) {
	uu := []entities.User{}
	err := faker.FakeData(
		&uu,
		options.WithRandomMapAndSliceMaxSize(10),
		options.WithRandomMapAndSliceMinSize(5),
	)
	if err != nil {
		return nil, err
	}

	return uu, nil
}

func GenerateOrders() ([]entities.Order, error) {

	oo := []entities.Order{}
	err := faker.FakeData(
		&oo,
		options.WithRandomMapAndSliceMaxSize(5),
		options.WithRandomMapAndSliceMinSize(1),
	)
	if err != nil {
		return nil, err
	}

	return oo, nil
}
