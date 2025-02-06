package handlers

import (
	"finance-control/models"
	"finance-control/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type InvestmentMovementHandler struct {
	repo repository.InvestmentMovementRepository
}

func NewInvestmentMovementHandler(repo repository.InvestmentMovementRepository) *InvestmentMovementHandler {
	return &InvestmentMovementHandler{repo: repo}
}

func (h *InvestmentMovementHandler) GetMovements(c *gin.Context) {
	movements, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching movements"})
		return
	}
	c.JSON(http.StatusOK, movements)
}

func (h *InvestmentMovementHandler) GetMovement(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	movement, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movement not found"})
		return
	}
	c.JSON(http.StatusOK, movement)
}

func (h *InvestmentMovementHandler) CreateMovement(c *gin.Context) {
	var input struct {
		InvestmentID uint      `json:"investment_id" binding:"required"`
		Amount       float64   `json:"amount" binding:"required"`
		MovementType string    `json:"movement_type" binding:"required,oneof=deposit withdrawal gain loss"`
		MovementDate time.Time `json:"movement_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movement := models.InvestmentMovement{
		InvestmentID: input.InvestmentID,
		Amount:       input.Amount,
		MovementType: input.MovementType,
		CreatedAt:    time.Now(),
		MovementDate: input.MovementDate,
	}
	if err := h.repo.Create(&movement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating movement"})
		return
	}
	c.JSON(http.StatusCreated, movement)
}

func (h *InvestmentMovementHandler) UpdateMovement(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input struct {
		InvestmentID uint      `json:"investment_id" binding:"required"`
		Amount       float64   `json:"amount" binding:"required"`
		MovementType string    `json:"movement_type" binding:"required,oneof=deposit withdrawal gain loss"`
		MovementDate time.Time `json:"movement_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movement, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movement not found"})
		return
	}
	movement.InvestmentID = input.InvestmentID
	movement.Amount = input.Amount
	movement.MovementType = input.MovementType
	movement.MovementDate = input.MovementDate
	if err := h.repo.Update(&movement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating movement"})
		return
	}
	c.JSON(http.StatusOK, movement)
}

func (h *InvestmentMovementHandler) DeleteMovement(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting movement"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movement deleted successfully"})
}
