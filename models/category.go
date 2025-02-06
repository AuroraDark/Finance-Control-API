package models

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required,oneof=income expense investment"`
}
