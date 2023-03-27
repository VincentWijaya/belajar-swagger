package models

import "time"

// Car represents the model for an car
type Car struct {
	ID        uint   `json:"id"`
	Model     string `json:"model"`
	Brand     string `json:"brand"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
