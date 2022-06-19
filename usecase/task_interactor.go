package usecase

import "github.com/moko-poi/go-api-todo/domain"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Add(u domain.Task) (err error) {
	err = interactor.TaskRepository.Store(u)
	return
}

func (interactor *TaskInteractor) Tasks() (task domain.Tasks, err error) {
	task, err = interactor.TaskRepository.FindAll()
	return
}

func (interactor *TaskInteractor) TaskById(identifier int) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.FindById(identifier)
	return
}
