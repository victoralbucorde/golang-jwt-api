package routes

import (
	"github.com/gin-gonic/gin"
	"ps-backend-victor-albuquerque-marcello-montella/testes/controllers"
	"ps-backend-victor-albuquerque-marcello-montella/testes/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/:userId", controllers.GetUser())
	incomingRoutes.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted"})
	})
	incomingRoutes.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted"})
	})
}
