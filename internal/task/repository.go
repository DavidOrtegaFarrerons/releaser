package task

type Repository interface {
	Create(task *Task) error
	GetByIdsAndType(ids []int, taskType TaskType) ([]Task, error)
	DeleteById(id int) error
}
