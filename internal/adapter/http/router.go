package http

import (
	"noveats-be/internal/adapter/http/handler"
	"noveats-be/internal/adapter/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(userHandler *handler.UserHandler, menuHandler *handler.MenuHandler, logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler(logger))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API V1 routes
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		menus := v1.Group("/menus")
		{
			// menus.POST("", menuHandler.CreateMenu)
			menus.GET("", menuHandler.GetAllMenu)
			menus.GET("/:id", menuHandler.GetMenu)
			// menus.PUT("/:id", menuHandler.UpdateMenu)
			// menus.DELETE("/:id", menuHandler.DeleteMenu)
		}
	}

	return router
}
