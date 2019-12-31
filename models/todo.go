package models

import "time"

type Todo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
	IsDeleted  bool   `json:"isDeleted"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
