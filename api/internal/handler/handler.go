package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateTask(c *gin.Context)
	GetTask(c *gin.Context)
}
