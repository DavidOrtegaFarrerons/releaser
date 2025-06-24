package task

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) AddTask(prId int, releaseId string, taskType TaskType, content string) error {
	task := &Task{
		PrId:      prId,
		ReleaseId: releaseId,
		Type:      taskType,
		Content:   content,
	}

	return s.Repo.Create(task)
}

func (s *Service) GetTasksByIdsAndType(prIds []int, taskType TaskType) ([]Task, error) {

	return s.Repo.GetByIdsAndType(prIds, taskType)
}
