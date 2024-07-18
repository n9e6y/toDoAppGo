package routes

import (
	"github.com/gin-gonic/gin"
	"todoapp/controllers"
)

func SetupRoutes(r *gin.Engine) {
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
