package model

import "time"

const (
	StatusPending    string = "pending"
	StatusInProgress string = "running"
	StatusCompleted  string = "success"
	StatusFailed     string = "failed"
)

type Task struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Payload    string `json:"payload"`
	Result     string `json:"result"`
	RetryCount int    `json:"retry_count"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
