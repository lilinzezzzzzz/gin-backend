package entity

import (
	"time"
)

type UserSessionData struct {
	Account   string     `json:"account"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	ID        uint       `json:"id"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Category  string     `json:"category"`
}

type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
