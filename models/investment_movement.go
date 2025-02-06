package models

import "time"

type InvestmentMovement struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	InvestmentID uint    `json:"investment_id" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	//"deposit", "withdrawal", "gain" or "loss"
	MovementType string    `json:"movement_type" binding:"required,oneof=deposit withdrawal gain loss"`
	CreatedAt    time.Time `json:"created_at"`
	MovementDate time.Time `json:"movement_date" binding:"required"`
}
