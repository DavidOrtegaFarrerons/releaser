package task

import "time"

type TaskType string

const (
	PreTask  TaskType = "PRE"
	PostTask TaskType = "POST"
)

// Task Entity
type Task struct {
	ID        int64
	PrId      string
	ReleaseId string
	Type      TaskType
	Content   string
	CreatedAt time.Time
}

type CreateTaskInput struct {
	PrId      string   `json:"prId"`
	ReleaseId string   `json:"releaseId"`
	Type      TaskType `json:"type"`
	Content   string   `json:"content"`
}
