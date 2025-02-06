package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	Password  string    `json:"-" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
