package models

import "time"

type Transaction struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	//"income" or "expense"
	Type            string    `json:"type"`
	TransactionDate time.Time `json:"transaction_date"`
	CategoryID      uint      `json:"category_id"`
	Category        Category  `json:"category"`
}
