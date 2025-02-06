package models

import "time"

type Investment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
