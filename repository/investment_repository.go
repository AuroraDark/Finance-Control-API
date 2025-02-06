package repository

import (
	"finance-control/models"

	"gorm.io/gorm"
)

// InvestmentRepository define os métodos de acesso aos dados para investimentos.
type InvestmentRepository interface {
	GetAll() ([]models.Investment, error)
	GetByID(id uint) (models.Investment, error)
	Create(investment *models.Investment) error
	Update(investment *models.Investment) error
	Delete(id uint) error
}

type investmentRepository struct {
	db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) InvestmentRepository {
	return &investmentRepository{db: db}
}

func (r *investmentRepository) GetAll() ([]models.Investment, error) {
	var investments []models.Investment
	err := r.db.Find(&investments).Error
	return investments, err
}

func (r *investmentRepository) GetByID(id uint) (models.Investment, error) {
	var investment models.Investment
	err := r.db.First(&investment, id).Error
	return investment, err
}

func (r *investmentRepository) Create(investment *models.Investment) error {
	return r.db.Create(investment).Error
}

func (r *investmentRepository) Update(investment *models.Investment) error {
	return r.db.Save(investment).Error
}

func (r *investmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Investment{}, id).Error
}
