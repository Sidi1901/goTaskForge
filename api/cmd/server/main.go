package main

import (
	"Sidi1901/goTaskForge/api/internal/config"
	"Sidi1901/goTaskForge/api/internal/database"
	"Sidi1901/goTaskForge/api/internal/handler"
	"Sidi1901/goTaskForge/api/internal/queue"
	"Sidi1901/goTaskForge/api/internal/repository"
	"Sidi1901/goTaskForge/api/internal/route"
	"Sidi1901/goTaskForge/api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load env if present

	cfg := config.LoadConfig()

	database.ConnectDB(cfg)
	queue.ConnectQueue(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Initialize services
	taskRepo := repository.NewTaskRepository(database.DB)
	taskSvc := service.NewTaskService(taskRepo)

	// Initialise Handlers
	taskHandler := handler.NewHandlerTask(taskSvc)

	// Initialize middlewares

	r := gin.New()

	route.SetupRoutes(r, taskHandler)

	r.Run(":" + cfg.ServerPort)

}
