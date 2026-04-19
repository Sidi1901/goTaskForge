package model

import "time"

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Status     TaskStatus `json:"status"`
	Payload    string     `json:"payload"`
	Result     string     `json:"result"`
	RetryCount int        `json:"retry_count"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
