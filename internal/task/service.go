package task

type Service struct {
	Repo Repository
}

func (s *Service) AddTask(prId string, releaseId string, taskType TaskType, content string) error {
	task := &Task{
		PrId:      prId,
		ReleaseId: releaseId,
		Type:      taskType,
		Content:   content,
	}

	return s.Repo.Create(task)
}
