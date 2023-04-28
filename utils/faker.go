package utils

import (
	"github.com/bxcodec/faker/v3"
)

// CreateFaker returns faker data with the type of T
func CreateFaker[T any]() (T, error) {
	// create a new variable called "fakerData"
	// to store the faker data
	var fakerData *T = new(T)

	// create a faker or sample data
	// the faker result will be stored to the "fakerData" variable
	err := faker.FakeData(fakerData)
	if err != nil {
		return *fakerData, err
	}

	// return the faker or sample data
	return *fakerData, nil
}
