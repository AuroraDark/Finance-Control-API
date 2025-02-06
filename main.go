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
	"finance-control/middleware"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Category{}, &models.Transaction{}, &models.User{}, &models.Investment{}, &models.InvestmentMovement{}); err != nil {
		log.Fatalf("Error in migration: %v", err)
	}

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

	router := gin.Default()
	store := cookie.NewStore([]byte("your_super_secret_key"))
	router.Use(sessions.Sessions("finance_session", store))

	api := router.Group("/api")
	{
		api.POST("/login", userHandler.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			// Auth and session
			protected.POST("/logout", userHandler.Logout)
			protected.GET("/session", userHandler.GetSessionUser)

			// Transactions
			protected.GET("/transactions", transactionHandler.GetTransactions)
			protected.POST("/transactions", transactionHandler.CreateTransaction)
			protected.PUT("/transactions/:id", transactionHandler.UpdateTransaction)
			protected.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

			// Categories
			protected.GET("/categories", categoryHandler.GetCategories)
			protected.GET("/categories/:id", categoryHandler.GetCategory)
			protected.POST("/categories", categoryHandler.CreateCategory)
			protected.PUT("/categories/:id", categoryHandler.UpdateCategory)
			protected.DELETE("/categories/:id", categoryHandler.DeleteCategory)

			// Users
			protected.GET("/users", userHandler.GetUsers)
			protected.GET("/users/:id", userHandler.GetUser)
			protected.POST("/users", userHandler.CreateUser)
			protected.PUT("/users/:id", userHandler.UpdateUser)
			protected.DELETE("/users/:id", userHandler.DeleteUser)

			// Investments
			protected.GET("/investments", investmentHandler.GetInvestments)
			protected.GET("/investments/:id", investmentHandler.GetInvestment)
			protected.POST("/investments", investmentHandler.CreateInvestment)
			protected.PUT("/investments/:id", investmentHandler.UpdateInvestment)
			protected.DELETE("/investments/:id", investmentHandler.DeleteInvestment)

			// Investment Movements
			protected.GET("/investment_movements", investmentMovementHandler.GetMovements)
			protected.GET("/investment_movements/:id", investmentMovementHandler.GetMovement)
			protected.POST("/investment_movements", investmentMovementHandler.CreateMovement)
			protected.PUT("/investment_movements/:id", investmentMovementHandler.UpdateMovement)
			protected.DELETE("/investment_movements/:id", investmentMovementHandler.DeleteMovement)
		}
	}

	// Route for Swagger documentation (if needed, uncomment later)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running on 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
