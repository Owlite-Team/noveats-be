package main

import (
	"database/sql"
	"fmt"
	"log"
	"noveats-be/config"
	"noveats-be/internal/adapter/http"
	"noveats-be/internal/adapter/http/handler"
	"noveats-be/internal/infrastructure/database"
	"noveats-be/internal/infrastructure/logger"
	"noveats-be/internal/repository/postgres"
	"noveats-be/internal/usecase/menu"
	"noveats-be/internal/usecase/user"

	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	zapLogger, err := logger.NewLogger(cfg.Log.Level)
	if err != nil {
		log.Fatal("Failed to initialize logger: ", err)
	}
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Fatal("Failed to sync logger: ", err)
		}
	}(zapLogger)

	zapLogger.Info("Starting application...")

	// Connect to database
	db, err := database.ConnectDb(cfg.Database)
	if err != nil {
		zapLogger.Fatal("Failed to connect to database: " + err.Error())
	}
	defer func(db *sql.DB) {
		err := database.CloseDB(db)
		if err != nil {
			zapLogger.Fatal("Failed to close database connection: " + err.Error())
		}
	}(db)

	// Initialize repositories
	userRepository := postgres.NewUserRepository(db)
	menuRepository := postgres.NewMenuRepository(db)

	// Initialize use cases
	createUserUC := user.NewCreateUserUseCase(userRepository)
	getUserUC := user.NewGetUserUseCase(userRepository)
	updateUserUC := user.NewUpdateUserUseCase(userRepository)
	deleteUserUC := user.NewDeleteUserUseCase(userRepository)
	createMenuUC := menu.NewCreateMenuUseCase(menuRepository)
	getMenuUC := menu.NewGetMenuUseCase(menuRepository)
	updateMenuUC := menu.NewUpdateMenuUseCase(menuRepository)
	deleteMenuUC := menu.NewDeleteMenuUseCase(menuRepository)

	// Initialize handlers
	userHandler := handler.NewUserHandler(
		createUserUC,
		getUserUC,
		updateUserUC,
		deleteUserUC,
		zapLogger,
	)
	menuHandler := handler.NewMenuHandler(
		createMenuUC,
		getMenuUC,
		updateMenuUC,
		deleteMenuUC,
		zapLogger,
	)

	// Setup router
	router := http.NewRouter(userHandler, menuHandler, zapLogger)

	// Start server
	address := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	zapLogger.Info("Server started on " + address)

	if err := router.Run(address); err != nil {
		zapLogger.Fatal("Failed to start server: " + err.Error())
	}
}
