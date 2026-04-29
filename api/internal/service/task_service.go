package service

import (
	"context"
	"encoding/json"

	"Sidi1901/goTaskForge/api/internal/dto"
	"Sidi1901/goTaskForge/api/internal/queue"
	"Sidi1901/goTaskForge/api/internal/repository"
	"Sidi1901/goTaskForge/shared/model"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, req dto.CreateTaskRequest) (string, error)
	GetTaskStatus(ctx context.Context, id string) (string, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (string, error) {
	payloadBytes, _ := json.Marshal(req.Data)

	task := &model.Task{
		ID:         uuid.New().String(),
		Status:     model.StatusPending,
		Payload:    string(payloadBytes),
		RetryCount: 0,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return "", err
	}

	if err := queue.EnqueueTask(ctx, task.ID); err != nil {
		return "", err
	}

	return task.ID, nil
}

func (s *taskService) GetTaskStatus(ctx context.Context, id string) (string, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return "", err
	}
	return task.Status, nil
}
