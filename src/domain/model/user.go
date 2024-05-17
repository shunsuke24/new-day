package model

import "time"

type User struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IconImageURL string    `json:"icon_image_url"`
	Name         string    `json:"name"`
}
