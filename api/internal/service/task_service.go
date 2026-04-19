package service

import (
	"Sidi1901/goTaskForge/api/internal/dto"
	"Sidi1901/goTaskForge/shared/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskService interface {
	CreateTask(req dto.CreateTaskRequest) (*model.Task, error)
	GetTask(id string) (*model.Task, error)
}

type taskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{db: db}
}

func (s *taskService) CreateTask(req dto.CreateTaskRequest) (*model.Task, error) {
	task := &model.Task{
		ID:         uuid.New().String(),
		Status:     model.StatusPending,
		Payload:    req.Payload,
		RetryCount: 0,
	}

	if err := s.db.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTask(id string) (*model.Task, error) {
	var task model.Task
	if err := s.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
