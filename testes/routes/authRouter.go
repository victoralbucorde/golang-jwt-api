package routes

import (
	"github.com/gin-gonic/gin"
	"ps-backend-victor-albuquerque-marcello-montella/testes/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/register", controllers.Register())
	incomingRoutes.POST("users/login", controllers.Login())
}
