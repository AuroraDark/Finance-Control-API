package handlers

import (
	"finance-control/models"
	"finance-control/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type InvestmentHandler struct {
	repo repository.InvestmentRepository
}

func NewInvestmentHandler(repo repository.InvestmentRepository) *InvestmentHandler {
	return &InvestmentHandler{repo: repo}
}

func (h *InvestmentHandler) GetInvestments(c *gin.Context) {
	investments, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching investments"})
		return
	}
	c.JSON(http.StatusOK, investments)
}

func (h *InvestmentHandler) GetInvestment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	investment, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Investment not found"})
		return
	}
	c.JSON(http.StatusOK, investment)
}

func (h *InvestmentHandler) CreateInvestment(c *gin.Context) {
	var input struct {
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	investment := models.Investment{
		Description: input.Description,
		CreatedAt:   time.Now(),
	}
	if err := h.repo.Create(&investment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating investment"})
		return
	}
	c.JSON(http.StatusCreated, investment)
}

func (h *InvestmentHandler) UpdateInvestment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input struct {
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	investment, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Investment not found"})
		return
	}
	investment.Description = input.Description
	if err := h.repo.Update(&investment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating investment"})
		return
	}
	c.JSON(http.StatusOK, investment)
}

func (h *InvestmentHandler) DeleteInvestment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting investment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Investment deleted successfully"})
}
