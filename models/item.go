package models

import (
	"time"
)

type Item struct {
	// the ID field will be filled with uuid data from the faker
	ID string `json:"id" faker:"uuid_hyphenated"`
	// the Name field will be filled with name data from the faker
	Name string `json:"name" faker:"name"`
	// the Price field will be filled with one of these values: 15, 27, 61
	Price int `json:"price" faker:"oneof: 15, 27, 61"`
	// the Quantity field will be filled with one of these values: 15, 27, 61
	Quantity  int       `json:"quantity" faker:"oneof: 15, 27, 61"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
