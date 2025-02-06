package repository

import (
	"finance-control/models"

	"gorm.io/gorm"
)

type InvestmentMovementRepository interface {
	GetAll() ([]models.InvestmentMovement, error)
	GetByID(id uint) (models.InvestmentMovement, error)
	Create(movement *models.InvestmentMovement) error
	Update(movement *models.InvestmentMovement) error
	Delete(id uint) error
}

type investmentMovementRepository struct {
	db *gorm.DB
}

func NewInvestmentMovementRepository(db *gorm.DB) InvestmentMovementRepository {
	return &investmentMovementRepository{db: db}
}

func (r *investmentMovementRepository) GetAll() ([]models.InvestmentMovement, error) {
	var movements []models.InvestmentMovement
	err := r.db.Find(&movements).Error
	return movements, err
}

func (r *investmentMovementRepository) GetByID(id uint) (models.InvestmentMovement, error) {
	var movement models.InvestmentMovement
	err := r.db.First(&movement, id).Error
	return movement, err
}

func (r *investmentMovementRepository) Create(movement *models.InvestmentMovement) error {
	return r.db.Create(movement).Error
}

func (r *investmentMovementRepository) Update(movement *models.InvestmentMovement) error {
	return r.db.Save(movement).Error
}

func (r *investmentMovementRepository) Delete(id uint) error {
	return r.db.Delete(&models.InvestmentMovement{}, id).Error
}
