// Author: Andressa Silva <developer.dressa@gmail.com>
// License: MIT
//
// This project is under development.
package main

import (
	"fmt"
	"log"

	"finance-control/config"
	"finance-control/handlers"
	"finance-control/models"
	"finance-control/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig("config/config.json")
	dbConfig := config.AppConfig.Database

	// Monta a string de conexão com PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode)

	// Conecta ao banco de dados utilizando GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migra os modelos: Category, Transaction, User, Investment e InvestmentMovement
	if err := db.AutoMigrate(&models.Category{}, &models.Transaction{}, &models.User{}, &models.Investment{}, &models.InvestmentMovement{}); err != nil {
		log.Fatalf("Error in migration: %v", err)
	}

	// Inicializa os repositórios e os handlers
	transactionRepo := repository.NewTransactionRepository(db)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	investmentRepo := repository.NewInvestmentRepository(db)
	investmentHandler := handlers.NewInvestmentHandler(investmentRepo)

	investmentMovementRepo := repository.NewInvestmentMovementRepository(db)
	investmentMovementHandler := handlers.NewInvestmentMovementHandler(investmentMovementRepo)

	// Configura o router do Gin e o middleware de sessão
	router := gin.Default()
	store := cookie.NewStore([]byte("your_super_secret_key"))
	router.Use(sessions.Sessions("finance_session", store))

	api := router.Group("/api")
	{
		// Transactions
		api.GET("/transactions", transactionHandler.GetTransactions)
		api.POST("/transactions", transactionHandler.CreateTransaction)
		api.PUT("/transactions/:id", transactionHandler.UpdateTransaction)
		api.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

		// Categories
		api.GET("/categories", categoryHandler.GetCategories)
		api.GET("/categories/:id", categoryHandler.GetCategory)
		api.POST("/categories", categoryHandler.CreateCategory)
		api.PUT("/categories/:id", categoryHandler.UpdateCategory)
		api.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		// Users
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		// Auth and session
		api.POST("/login", userHandler.Login)
		api.POST("/logout", userHandler.Logout)
		api.GET("/session", userHandler.GetSessionUser)

		// Investments
		api.GET("/investments", investmentHandler.GetInvestments)
		api.GET("/investments/:id", investmentHandler.GetInvestment)
		api.POST("/investments", investmentHandler.CreateInvestment)
		api.PUT("/investments/:id", investmentHandler.UpdateInvestment)
		api.DELETE("/investments/:id", investmentHandler.DeleteInvestment)

		// Invesment movements
		api.GET("/investment_movements", investmentMovementHandler.GetMovements)
		api.GET("/investment_movements/:id", investmentMovementHandler.GetMovement)
		api.POST("/investment_movements", investmentMovementHandler.CreateMovement)
		api.PUT("/investment_movements/:id", investmentMovementHandler.UpdateMovement)
		api.DELETE("/investment_movements/:id", investmentMovementHandler.DeleteMovement)
	}

	// Route to docs in swagger (under development)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running on 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
