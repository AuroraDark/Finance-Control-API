package handlers

import (
	"finance-control/models"
	"finance-control/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	var startTime *time.Time
	var endTime *time.Time

	if startParam := c.Query("start"); startParam != "" {
		t, err := time.Parse("2006-01-02", startParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for 'start'. Use YYYY-MM-DD."})
			return
		}
		startTime = &t
	}

	if endParam := c.Query("end"); endParam != "" {
		t, err := time.Parse("2006-01-02", endParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for 'end'. Use YYYY-MM-DD."})
			return
		}
		endTime = &t
	}

	// Uses the transaction_date column for filtering
	transactions, err := h.repo.GetAllWithFilters(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input struct {
		Description     string    `json:"description" binding:"required"`
		Amount          float64   `json:"amount" binding:"required"`
		Type            string    `json:"type" binding:"required,oneof=income expense"`
		CategoryID      uint      `json:"category_id" binding:"required"`
		TransactionDate time.Time `json:"transaction_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		Description:     input.Description,
		Amount:          input.Amount,
		Type:            input.Type,
		TransactionDate: input.TransactionDate,
		CategoryID:      input.CategoryID,
	}

	if err := h.repo.Create(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Description string  `json:"description" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
		// Accepts only "income" or "expense"
		Type            string    `json:"type" binding:"required,oneof=income expense"`
		CategoryID      uint      `json:"category_id" binding:"required"`
		TransactionDate time.Time `json:"transaction_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		ID:              uint(id),
		Description:     input.Description,
		Amount:          input.Amount,
		Type:            input.Type,
		CategoryID:      input.CategoryID,
		TransactionDate: input.TransactionDate,
	}

	if err := h.repo.Update(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
