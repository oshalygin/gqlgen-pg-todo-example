package models

import "time"

type Todo struct {
	ID         int    `json:",pk,unique,notnull"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
	IsDeleted  bool   `json:"isDeleted"`

	CreatedBy int `json:"createdBy" pg:"fk:user_id"`
	UpdatedBy int `json:"updatedBy" pg:"fk:user_id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
