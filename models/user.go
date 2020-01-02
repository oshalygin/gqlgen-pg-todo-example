package models

import "time"

type User struct {
	ID    int    `json:"id" pg:",pk,unique,notnull"`
	Email string `json:"email" pg:",unique,notnull"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
