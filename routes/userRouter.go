package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitkoomar/JTWauthentication/controllers"
	"github.com/sumitkoomar/JTWauthentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
