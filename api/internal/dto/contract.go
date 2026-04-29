package dto

type CreateTaskRequest struct {
	Type string      `json:"type" binding:"required"`
	Data interface{} `json:"data"`
}
