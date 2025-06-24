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
	PrId      int
	ReleaseId string
	Type      TaskType
	Content   string
	CreatedAt time.Time
}

type CreateTaskInput struct {
	PrId      int      `json:"prId"`
	ReleaseId string   `json:"releaseId"`
	Type      TaskType `json:"type"`
	Content   string   `json:"content"`
}

type GetTasksByIdInput struct {
	PrIds []int    `json:"prIds"`
	Type  TaskType `json:"type"`
}
