package usecase

import "github.com/moko-poi/go-api-todo/domain"

type TaskRepository interface {
	Store(domain.Task) error
	FindById(int) (domain.Task, error)
	FindAll() (domain.Tasks, error)
}
