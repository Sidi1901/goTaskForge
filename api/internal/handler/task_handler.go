package handler

import (
	"errors"
	"net/http"

	"Sidi1901/goTaskForge/api/internal/dto"
	"Sidi1901/goTaskForge/api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlerTask struct {
	svc service.TaskService
}

func NewHandlerTask(svc service.TaskService) Handler {
	return &handlerTask{svc: svc}
}

func (h *handlerTask) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.svc.CreateTask(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *handlerTask) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.svc.GetTask(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
