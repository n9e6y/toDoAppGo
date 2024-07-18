package routes

import (
	"github.com/gin-gonic/gin"
	"todoapp/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
	}
}
