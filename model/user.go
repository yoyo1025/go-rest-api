package model

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`  // 重複を許さないユニークなフィールドになる
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}