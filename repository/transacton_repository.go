package repository

import (
	"finance-control/models"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllWithFilters(start, end *time.Time) ([]models.Transaction, error)
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetAllWithFilters(start, end *time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := r.db.Preload("Category")
	if start != nil {
		query = query.Where("transaction_date >= ?", *start)
	}
	if end != nil {
		query = query.Where("transaction_date <= ?", *end)
	}
	err := query.Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}
