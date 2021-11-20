package models

type Good struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	CategoryID string `json:"category_id" db:"category_id"`
}
