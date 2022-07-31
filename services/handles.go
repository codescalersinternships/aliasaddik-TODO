package services

import (
	"github.com/aliasaddik/todo-project/models"
)

type Handles interface {
	CreateTask(*models.Task) error
	GetTask() ([]*models.Task, error)
	EditTask(*models.Task) error
	DeleteTask(string) error
}
