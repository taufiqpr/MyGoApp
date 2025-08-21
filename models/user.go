package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:50;not null"`
	Email        string    `gorm:"uniqueIndex;size:100;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=5,max=50"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=5,max=50"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UpdateProfileRequest struct {
    Name     string `json:"name" binding:"omitempty,min=3,max=50"`
    Email    string `json:"email" binding:"omitempty,email,max=100"`
    Password string `json:"password" binding:"omitempty,min=5,max=50"`
}
