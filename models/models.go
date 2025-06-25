package models

import "time"

type Contact struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactInput struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Phone string `form:"phone" json:"phone"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
