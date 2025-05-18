package task

type Repository interface {
	Create(task *Task) error
	ListByReleaseId(releaseId string) ([]Task, error)
	DeleteById(id int) error
}
