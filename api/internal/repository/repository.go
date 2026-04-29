package repository

import (
	"Sidi1901/goTaskForge/shared/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *model.Task) error
	GetTaskByID(id string) (*model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTaskByID(id string) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
