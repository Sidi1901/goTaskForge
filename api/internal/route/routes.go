package route

import (
	"Sidi1901/goTaskForge/api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, taskHandler handler.Handler) {

	api := r.Group("/api/v1")
	{
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks/:id", taskHandler.GetTask)
	}
}
