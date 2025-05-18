package task

import "time"

type TaskType string

const (
	PreTask  TaskType = "PRE"
	PostTask TaskType = "POST"
)

type Task struct {
	ID        int64
	PrId      string
	ReleaseId string
	Type      TaskType
	Content   string
	CreatedAt time.Time
}
