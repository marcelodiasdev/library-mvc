package models

import "time"

type User struct {
	ID        int64     `json:"ID" binding:"required,min=10,max=200"`
	Name      string    `json:"name" binding:"required,email"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
