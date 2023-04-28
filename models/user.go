package models

import "time"

type User struct {
	// the ID field will be filled with uuid data from the faker
	ID string `json:"id" faker:"uuid_hyphenated"`
	// the Email field will be filled with email data from the faker
	Email string `json:"email" gorm:"unique" faker:"email"`
	// the Password field will be filled with password data from the faker
	Password  string    `json:"password" faker:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
